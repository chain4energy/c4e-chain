package keeper

import (
	"time"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateCampaign(ctx sdk.Context, owner string, name string, description string, campaignType types.CampaignType, removableClaimRecords bool,
	feeGrantAmount *math.Int, initialClaimFreeAmount *math.Int, free *sdk.Dec, startTime *time.Time,
	endTime *time.Time, lockupPeriod *time.Duration, vestingPeriod *time.Duration, vestingPoolName string) (*types.Campaign, error) { // TODO za duzo tych pointerow, odpointerowanie powinno byc na poziomie obslugi message w msg server
	k.Logger(ctx).Debug("create campaign", "owner", owner, "name", name, "description", description,
		"startTime", startTime, "endTime", endTime, "lockupPeriod", lockupPeriod, "vestingPeriod", vestingPeriod)

	validFree, err := validateFreeAmount(free) // TODO pwoinno byc w types.ValidateCreateCampaignParams
	if err != nil {
		return nil, err
	}

	if err = k.ValidateCampaignParams(ctx, name, description, validFree, startTime, endTime, campaignType, owner, vestingPoolName, lockupPeriod, vestingPeriod); err != nil {
		return nil, err
	}
	if err = ValidateCampaignEndTimeInTheFuture(ctx, endTime); err != nil {
		return nil, err
	}
	validdFeegrantAmount, err := validateFeegrantAmount(feeGrantAmount) // TODO pwoinno byc w types.ValidateCreateCampaignParams
	if err != nil {
		return nil, err
	}
	validInitialClaimFreeAmount, err := validateInitialClaimFreeAmount(initialClaimFreeAmount)
	if err != nil {
		return nil, err
	}

	campaign := types.Campaign{
		Owner:                  owner,
		Name:                   name,
		Description:            description,
		CampaignType:           campaignType,
		RemovableClaimRecords:  removableClaimRecords,
		FeegrantAmount:         validdFeegrantAmount,
		InitialClaimFreeAmount: validInitialClaimFreeAmount,
		Free:                   validFree,
		Enabled:                false,
		StartTime:              *startTime,
		EndTime:                *endTime,
		LockupPeriod:           *lockupPeriod,
		VestingPeriod:          *vestingPeriod,
		VestingPoolName:        vestingPoolName,
	}

	campaignId := k.AppendNewCampaign(ctx, campaign)

	missionInitial := types.NewInitialMission(campaignId) // TODO troche dziwne ze tworzymy obiekt mission a potem tylko przekazujemy jego paramtry, tutaj jedyne miejsce, dlczego poprsu nie przekazac pramtwrow oczekiwachych? po co ten dodatkowy krok?
	err = k.AddMissionToCampaign(ctx, owner, campaignId, missionInitial.Name, missionInitial.Description,
		missionInitial.MissionType, missionInitial.Weight, missionInitial.ClaimStartDate)
	if err != nil {
		return nil, err
	}

	return &campaign, nil
}

func validateInitialClaimFreeAmount(initialClaimFreeAmount *math.Int) (math.Int, error) { // TODO nazwa co w stylu fixAndValidate i do types/campaign.go
	if initialClaimFreeAmount == nil {  // TODO Te 2 nil na 0 jest z tego co mi sie wydaje w wielu miejscach w roznych modulach, wiec jakis common util by sie przydal NilIntToZeroInt
		return math.ZeroInt(), nil
	}
	if initialClaimFreeAmount.IsNil() {
		return math.ZeroInt(), nil
	}

	if initialClaimFreeAmount.IsNegative() { // TODO to trzba jak fukcje validcji w types/campaign.go i tutaj wywolane
		return math.ZeroInt(), errors.Wrapf(c4eerrors.ErrParam, "initial claim free amount (%s) cannot be negative", initialClaimFreeAmount.String())
	}

	return *initialClaimFreeAmount, nil
}

func validateFreeAmount(free *sdk.Dec) (sdk.Dec, error) { // TODO nazwa co w stylu fixAndValidate i do types/campaign.go ale racze nie bedzie potrban mo prniesieniu nil to 0 do obslugi Msg
	if free == nil {   // TODO Te 2 nil na 0 jest z tego co mi sie wydaje w wielu miejscach w roznych modulach, wiec jakis common util by sie przydal NilDecToZeroDec
		return sdk.ZeroDec(), nil
	}
	if free.IsNil() {
		return sdk.ZeroDec(), nil
	}

	if free.IsNegative() {
		return sdk.ZeroDec(), errors.Wrapf(c4eerrors.ErrParam, "free amount (%s) cannot be negative", free.String()) // TODO to nir etylko negative ale i <= 100 - to trzba jak fukcje validcji w types/campaign.go i tutaj wywolane
	}

	return *free, nil
}

func (k Keeper) CloseCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	k.Logger(ctx).Debug("close campaign", "owner", owner, "campaignId", campaignId)
	campaign, err := k.ValidateCloseCampaignParams(ctx, campaignId, owner)
	if err != nil {
		return err
	}
	campaign.Enabled = false
	k.SetCampaign(ctx, campaign) // TODO czy ten set campaign jest tutaj potrzebny? raczej na koniec po wukonaniu wsztkich operacji
	// TODO ========= ten caly blok jest duplikowany w Remove campaign - mozna osbna funkcje zrobic - returnAllToOwner lub cos podobnego
	if err = k.sendCampaignCurrentAmountToOwner(ctx, &campaign, campaign.CampaignCurrentAmount); err != nil { 
		return err
	}
	if err = k.closeCampaignSendFeegrant(ctx, &campaign); err != nil {
		return err
	}
//	===========================================
	return nil
}

func (k Keeper) sendCampaignCurrentAmountToOwner(ctx sdk.Context, campaign *types.Campaign, amount sdk.Coins) error { // TODO nazwa co w stylu returnToOwner
	if amount.IsAnyGT(campaign.CampaignCurrentAmount) { // TODO to prosba o dobre przetestowanie unitowe bo te gupowe porownywania to w jakis przypadkach dzialaly niespodziwanie, tylko juz nie pamotam co konkretnie bylo i czy to ta funckja (cos mi swuta o przypadki gdy jakis token istnije tylko w jednej liscie)
		return errors.Wrapf(c4eerrors.ErrAmount,
			"cannot send campaign current amount to owner, campaign current amount is lower than amount (%s < %s)", campaign.CampaignCurrentAmount, amount)
	}
	if campaign.CampaignType == types.VestingPoolCampaign {
		if err := k.vestingKeeper.RemoveVestingPoolReservation(ctx, campaign.Owner, campaign.VestingPoolName, campaign.Id,
			amount.AmountOf(k.vestingKeeper.Denom(ctx))); err != nil {
			return err
		}
	} else {
		ownerAddress, _ := sdk.AccAddressFromBech32(campaign.Owner)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddress, amount); err != nil {
			return err
		}
	}

	campaign.CampaignCurrentAmount = campaign.CampaignCurrentAmount.Sub(amount...)
	k.SetCampaign(ctx, *campaign) // TODO cos z tym setem trzba omysles bo widac ze sa potem wielokrotne zmiany ustwiwane na kv store
	// genralnie pwonnismy zorbic zasada dla metod prywantnych ze jak campaign jest przekazywana do metody to tam metoda co najwyzej ja zwraca, napisuje wtedy metoda nardzedna 
	return nil
}

func (k Keeper) EnableCampaign(ctx sdk.Context, owner string, campaignId uint64, startTime *time.Time, endTime *time.Time) error {
	k.Logger(ctx).Debug("start campaign", "owner", owner, "campaignId", campaignId)

	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return err
	}

	if startTime != nil {
		campaign.StartTime = *startTime
	}
	if endTime != nil {
		campaign.EndTime = *endTime
	}

	err = k.ValidateEnableCampaignParams(ctx, campaign, owner)
	if err != nil {
		return err
	}

	campaign.Enabled = true
	k.SetCampaign(ctx, campaign)
	return nil
}

func (k Keeper) RemoveCampaign(ctx sdk.Context, owner string, campaignId uint64) error {
	k.Logger(ctx).Debug("remove campaign", "owner", owner, "campaignId", campaignId)

	campaign, err := k.ValidateRemoveCampaignParams(ctx, owner, campaignId)
	if err != nil {
		k.Logger(ctx).Debug("remove campaign", "err", err.Error())
		return err
	}
	// TODO ========= ten caly blok jest duplikowany w close campaign - mozna osbna funkcje zrobic - returnAllToOwner lub cos podobnego

	if err = k.sendCampaignCurrentAmountToOwner(ctx, campaign, campaign.CampaignCurrentAmount); err != nil {
		return err
	}
	if err = k.closeCampaignSendFeegrant(ctx, campaign); err != nil {
		return err
	}
// =================================================================
	k.removeCampaign(ctx, campaignId)
	k.RemoveAllMissionForCampaign(ctx, campaignId)
	return nil
}

// TODO czy te wszystkie validacje nie powinny byc prywatne?
// zasada validacji powinna byc taka ze te lementy ktore nie wymagaja Ctx pwinny byc walidowane w types
// te z ctx w keeperze

func (k Keeper) ValidateCampaignParams(ctx sdk.Context, name string, description string, free sdk.Dec, startTime *time.Time, endTime *time.Time, // TODO tutaj pownien byc start time i end time nie pointerem, odpointerowanie w obsludze Msg bo tam jest problem, po co go tutaj przenosic
	campaignType types.CampaignType, owner string, vestingPoolName string, lockupPeriod *time.Duration, vestingPeriod *time.Duration) error {
	if err := types.ValidateCreateCampaignParams(name, description, startTime, endTime, campaignType, vestingPoolName); err != nil {
		return err
	}

	if campaignType == types.VestingPoolCampaign {
		return k.ValidateCampaignWhenAddedFromVestingPool(ctx, owner, vestingPoolName, lockupPeriod, vestingPeriod, free)
	}
	return nil
}
func (k Keeper) ValidateCloseCampaignParams(ctx sdk.Context, campaignId uint64, owner string) (types.Campaign, error) {
	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return types.Campaign{}, err
	}
	if err = ValidateOwner(campaign, owner); err != nil {
		return types.Campaign{}, err
	}
	if err = ValidateCampaignEnded(ctx, campaign); err != nil {
		return types.Campaign{}, err
	}

	return campaign, nil
}

func (k Keeper) ValidateEnableCampaignParams(ctx sdk.Context, campaign types.Campaign, owner string) error {
	if err := ValidateOwner(campaign, owner); err != nil {
		return err
	}

	if err := types.ValidateCampaignIsNotEnabled(campaign); err != nil {
		return err
	}
	if err := types.ValidateCampaignEndTimeAfterStartTime(&campaign.StartTime, &campaign.EndTime); err != nil {
		return err
	}
	return ValidateCampaignEndTimeInTheFuture(ctx, &campaign.EndTime)
}

func (k Keeper) ValidateRemoveCampaignParams(ctx sdk.Context, owner string, campaignId uint64) (*types.Campaign, error) {
	campaign, err := k.ValidateCampaignExists(ctx, campaignId)
	if err != nil {
		return nil, err
	}

	if err = ValidateOwner(campaign, owner); err != nil {
		return nil, err
	}

	return &campaign, types.ValidateCampaignIsNotEnabled(campaign)
}

func (k Keeper) ValidateCampaignExists(ctx sdk.Context, campaignId uint64) (types.Campaign, error) { // TODO nazwa MustGetCampaign i do campaign_store.go
	campaign, found := k.GetCampaign(ctx, campaignId)
	if !found {
		return types.Campaign{}, errors.Wrapf(c4eerrors.ErrNotExists, "campaign with id %d not found", campaignId)
	}
	return campaign, nil
}

func ValidateOwner(campaign types.Campaign, owner string) error { // TODO to powinna byc metoda struktury Campaign. Zasada jak przekazuje caly obiek campaign to przesunac do structury Campaign
	if campaign.Owner != owner {
		return errors.Wrap(c4eerrors.ErrWrongSigner, "you are not the campaign owner") // TODO raczej address %s in not an owner
	}
	return nil
}

func ValidateCampaignEndTimeInTheFuture(ctx sdk.Context, endTime *time.Time) error { // TODO to powinna byc metoda struktury Campaign w postaci - validateNotEnded z parametrem Time bez pointera. Zasada jak przekazuje caly obiek campaign to przesunac do structury Campaign
	if endTime == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "create claim campaign - start time is nil error")
	}
	if endTime.Before(ctx.BlockTime()) { 
		return errors.Wrapf(c4eerrors.ErrParam, "end time in the past error (%s < %s)", endTime, ctx.BlockTime())
	}
	return nil
}

func ValidateCampaignNotEnded(ctx sdk.Context, campaign types.Campaign) error { // TODO to powinna byc metoda struktury Campaign w postaci - validateNotEnded z parametrem Time z parametrem Time bez pointera. Zasada jak przekazuje zaly obiek campaign to przesunac do structury Campaign. Generalnie ta metoda wyglda tak samo jak ValidateCampaignEndTimeInTheFuture
	if ctx.BlockTime().After(campaign.EndTime) {
		return errors.Wrapf(c4eerrors.ErrParam, "campaign with id %d campaign is over (end time - %s < %s)", campaign.Id, campaign.EndTime, ctx.BlockTime())
	}
	return nil
}

func ValidateCampaignEnded(ctx sdk.Context, campaign types.Campaign) error { // TODO to powinna byc metoda struktury Campaign w postaci - validateEnded z parametrem Time bez pointera.
	if ctx.BlockTime().Before(campaign.EndTime) { // TODO or equal
		return errors.Wrapf(c4eerrors.ErrParam, "campaign with id %d campaign is not over yet (endtime - %s < %s)", campaign.Id, campaign.EndTime, ctx.BlockTime())
	}
	return nil
}
