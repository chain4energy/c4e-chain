All changes must be comitted
```bash
git stash
```

```bash
git checkout v1.0.0
```
```bash
ignite chain build
```

Start chain with force delete data

```bash
ignite chain serve -f --path  $DIR -v $@ 2>$DIR/.log/log.txt &
```

Wait unit blochain start producing block
```bash
ignite chain serve -f --path  $DIR -v $@ 2>$DIR/.log/log.txt &
```

Prepare proposal from alice account
```bash
c4ed tx gov submit-proposal software-upgrade v2 --upgrade-height 75 --from alice --title "Upgrade to v2" --description "ASdA" --chain-id c4echain```
```

Deposit to proposal
```bash
c4ed tx gov deposit 1 100000000000uc4e --from alice --chain-id c4echain
```

Vote for proposal
```bash
c4ed tx gov vote 1 yes --from alice --chain-id c4echain
```

Wait until blochain halt on 75 height

Stash changes
```bash
git stash
```

Switch to newest blochain version
```bash
git switch -
```

Build project once again
```bash
ignite chain build
```

Start chain with actual data without force

```bash
ignite chain serve --path  $DIR -v $@ 2>$DIR/.log/log.txt &
```

All logs from migration are stored .log/log.txt &