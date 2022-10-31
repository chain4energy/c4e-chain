# How to test if migration works

---

1. Stash all changes in the current version

```bash
git stash
```

2. Checkout to older version

```
git checkout v1.0.0
```

3. Create a log folder if it does not exist

```
mkdir .log/
```

4. Build and serve chain with reset state option

```bash
ignite chain build
ignite chain serve -f --path <appDir> -v $@ 2>><appDir>/.log/log.txt
```

5. Wait unit blochain start producing block

---

6. Prepare proposal from alice account

```bash
c4ed tx gov submit-proposal software-upgrade <migrationName> --upgrade-height <upgradeHeight> --title <title> --description <description> --from <accName> --chain-id <chainId>
```

7. Deposit to proposal

```bash
c4ed tx gov deposit 1 100000000000uc4e --from alice --chain-id c4echain
```

8. Vote for proposal

```bash
c4ed tx gov vote 1 yes --from alice --chain-id c4echain
```

9. Wait until blochain halt on 75 height and stop the chain

---

10. Stash changes

```bash
git stash
```

11. Switch to newest blochain version

```bash
git switch -
```
12. Build project once again

```bash
ignite chain build
```

13. Copy app state (backup)

```bash
cp -r .data/ .backup_data_block_74/ 
```

14. Start chain with actual data without force

```bash
ignite chain serve --path <appDir> -v $@ 2>><appDir>/.log/log.txt &
```
All logs from migration are stored .log/log.txt &

## Troubleshoting 

If, after starting the new version of the application, ignite displays information about resetting the application state, 
stop the chain and copy the previously saved backup to data

```bash
rm -rf .data/
cp -r .backup_data_block_74/ .data/
```

