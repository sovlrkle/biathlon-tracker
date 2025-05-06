# System prototype for biathlon competitions

## Configuration (json)

- **Laps**        - Amount of laps for main distance
- **LapLen**      - Length of each main lap
- **PenaltyLen**  - Length of each penalty lap
- **FiringLines** - Number of firing lines per lap
- **Start**       - Planned start time for the first competitor
- **StartDelta**  - Planned interval between starts

## Run program

The program may be run with the following command:

```bash
env CONFIG_PATH=<?> EVENTS_PATH=<?> go run ./cmd/app/main.go
```

where `CONFIG_PATH` and `EVENTS_PATH` are environment variables:

- `CONFIG_PATH` - path to the file with events
- `EVENTS_PATH` - path to the file with configuration

The program will write a result table to `report` file in the project root directory.
If it already exists, it will be overwritten. Output logs will be written to `output.log` file in the project root
directory.
If it already exists, it will be overwritten.

## Examples

1. `config.json`

```json
{
  "laps": 2,
  "lapLen": 3651,
  "penaltyLen": 50,
  "firingLines": 1,
  "start": "09:30:00.000",
  "startDelta": "00:00:30"
}
```

`events`

```
[09:05:59.867] 1 1
[09:15:00.841] 2 1 09:30:00.000
[09:29:45.734] 3 1
[09:30:01.005] 4 1
[09:49:31.659] 5 1 1
[09:49:33.123] 6 1 1
[09:49:34.650] 6 1 2
[09:49:35.937] 6 1 4
[09:49:37.364] 6 1 5
[09:49:38.339] 7 1
[09:49:55.915] 8 1
[09:51:48.391] 9 1
[09:59:03.872] 10 1
[09:59:03.872] 11 1 Lost in the forest
```

`output.log`

```
[09:05:59.867] The competitor(1) registered
[09:15:00.841] The start time for the competitor(1) was set by a draw to 09:30:00.000
[09:29:45.734] The competitor(1) is on the start line
[09:30:01.005] The competitor(1) has started
[09:49:31.659] The competitor(1) is on the firing range(1)
[09:49:33.123] The target(1) has been hit by competitor(1)
[09:49:34.650] The target(2) has been hit by competitor(1)
[09:49:35.937] The target(4) has been hit by competitor(1)
[09:49:37.364] The target(5) has been hit by competitor(1)
[09:49:38.339] The competitor(1) left the firing range
[09:49:55.915] The competitor(1) entered the penalty laps
[09:51:48.391] The competitor(1) left the penalty laps
[09:59:03.872] The competitor(1) ended the main lap
[09:59:05.872] The competitor(1) can`t continue: Lost in the forest
```

`report`

```
[NotFinished] 1 [{00:29:02.867, 2.095}, {,}] {00:01:52.476, 0.445} 4/5
```

2. example from `.zip` with the task

`report`

```
[00:00:01.744] 1 [{00:12:33.636, 4.644}, {00:12:50.667, 4.542}] {00:02:30.000, 3.000} 7/10
[00:00:01.503] 2 [{00:12:38.243, 4.616}, {00:12:38.610, 4.614}] {00:01:40.000, 3.000} 8/10
[00:00:00.887] 3 [{00:12:42.386, 4.591}, {00:12:51.500, 4.537}] {,} 10/10
[00:00:01.278] 4 [{00:12:45.669, 4.571}, {00:13:19.466, 4.378}] {00:01:40.000, 3.000} 8/10
[00:00:00.331] 5 [{00:13:20.939, 4.370}, {00:13:01.202, 4.480}] {00:02:30.000, 3.000} 7/10

```