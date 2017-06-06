# Raspberry Pi3

This is a collection of generic steps you probably want to take with your Raspberry Pi3.

## Realtek 8192cu WiFi Chip

### Enable WiFi service interrupt tolerance
TODO: Add steps to allow a Raspberry Pi to tolerate WiFi service interrups.

### Disable sleep/suspend/hibernate

First, check the built-in/current power management setting for your chip:
```
$ cat /sys/module/8192cu/parameters/rtw_power_mgnt
```

A value greater than 0 means that power management is active for your WiFi chip, to disable it:
```
$ vi /etc/modprobe.d/8192cu.conf
```
and add the following:
```
# Disable built-in 8192cu power management settings.
options 8192cu rtw_power_mgnt=0
```
