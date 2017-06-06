# Raspberry Pi3

This is a collection of generic steps you probably want to take with your Raspberry Pi3.

### Disable power management on your Realtek 8192cu WiFi chip.

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
