## Furya-1.6.0 Upgrade Instructions

### Software Version and Key Dates

We will be upgrading node from v1.5.8 to v1.6.0.
The network will be shutdown with a SoftwareUpgradeProposal that activates at block height <height template>, which is approximately <date template>.
The version of cosmos-sdk becomes v0.45.9
The version of tendermint becomes is v0.34.21
The recommended version of golang is **1.19**

### Risks

As a validator, performing the upgrade procedure on your consensus nodes carries a risk of being slashed in case of not running the node in time.  
If you discover a mistake in the process, the best thing to do is to seek advice from a Furya developer before resetting your validator.

### Recovery

Prior to the upgrading procedure, validators are encouraged to take a full data snapshot at the export height before proceeding. Snap-shotting depends heavily on infrastructure, but generally this can be done by backing up the .furya and .furyacli directories.

In the event that the upgrade does not succeed, validators and operators must restore their nodes from backup with v1.5.8 of the furya software. Before starting the node validators should add `--unsafe-skip-upgrades <height template>` to as furya node start parameter.

### Upgrade Procedure

#### Before the upgrade

Furya has submitted a SoftwareUpgradeProposal that specifies block height <height template> as the final block height for <date template>. This height corresponds to approximately <date template>. Once the proposal passes, the chain will shutdown automatically at the specified height and does not require manual intervention by validators.

#### On the day of the upgrade

The furya chain is expected to halt at block height <height template>, at approximately <date template>, and restart with new software after an hour at <date template + hour>. Do not stop your node and begin the upgrade before <date template>, or you may go offline and be unable to recover until after the upgrade!

Make sure the furya process is stopped before proceeding and that you have backed up your validator. Failure to backup your validator could make it impossible to restart your node if the upgrade fails.

#### Guide

1. Stop the service that's running the node
```shell
sudo systemctl stop furya_node
```

2. Backup .furya and .furyacli
```shell
cp -rf $HOME/.furya $HOME/.furya.bak
cp -rf $HOME/.furyacli $HOME/.furyacli.bak
```

3. Check if you have proper go version
```shell
go version
```
It has to be `1.19` or higher. If it's not you should [install go with appropriate version](https://go.dev/doc/install).

4. Clone Furya from the latest release
```shell
git clone -b v1.6.0 https://github.com/TessorNetwork/furya
cd furya
```
  
5. Compile and install new version of Furya
```shell
make install
```
and check version
  
```shell
furya version
```
It has to be `1.6.0`

6. Start your node back
```shell
sudo systemctl start furya_node
```

7. Validate your node is up
```shell
sudo journalctl -u furya_node.service -f
```

## Coordination

If the network does not launch by <date template + 6 hours>, the launch should be considered a failure and validators should refer to the rollback instructions to restart the previous version of software. In the event of launch failure, coordination will occur in the Furya discord.
