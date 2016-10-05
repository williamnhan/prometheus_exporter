package main

type Info struct {
  PassengerVersion        string       `xml:"passenger_version"`
  AppCount                string       `xml:"group_count"`
  CurrentProcessCount     string       `xml:"process_count"`
  MaxProcessCount         string       `xml:"max"`
  CapacityUsed            string       `xml:"capacity_used"`
  TopLevelRequestsInQueue string       `xml:"get_wait_list_size"`
  SuperGroups             []SuperGroup `xml:"supergroups>supergroup"`
}

type SuperGroup struct {
  Name            string `xml:"name"`
  State           string `xml:"state"`
  RequestsInQueue string `xml:"get_wait_list_size"`
  CapacityUsed    string `xml:"capacity_used"`
  Group           Group  `xml:"group"`
}

type Group struct {
  Default               string    `xml:"default,attr"`
  Name                  string    `xml:"name"`
  ComponentName         string    `xml:"component_name"`
  AppRoot               string    `xml:"app_root"`
  AppType               string    `xml:"app_type"`
  Environment           string    `xml:"environment"`
  UUID                  string    `xml:"uuid"`
  EnabledProcessCount   string    `xml:"enabled_process_count"`
  DisablingProcessCount string    `xml:"disabling_process_count"`
  DisabledProcessCount  string    `xml:"disabled_process_count"`
  CapacityUsed          string    `xml:"capacity_used"`
  GetWaitListSize       string    `xml:"get_wait_list_size"`
  DisableWaitListSize   string    `xml:"disable_wait_list_size"`
  ProcessesSpawning     string    `xml:"processes_being_spawned"`
  LifeStatus            string    `xml:"life_status"`
  User                  string    `xml:"user"`
  UID                   string    `xml:"uid"`
  GID                   string    `xml:"gid"`
  Options               Options   `xml:"options"`
  Processes             []Process `xml:"processes>process"`
}

type Options struct {
  AppRoot                   string `xml:"app_root"`
  AppGroupName              string `xml:"app_group_name"`
  AppType                   string `xml:"app_type"`
  StartCommand              string `xml:"start_command"`
  StartupFile               string `xml:"startup_file"`
  ProcessTitle              string `xml:"process_title"`
  LogLevel                  string `xml:"log_level"`
  StartTimeout              string `xml:"start_timeout"`
  Environment               string `xml:"environment"`
  BaseURI                   string `xml:"base_uri"`
  SpawnMethod               string `xml:"spawn_method"`
  DefaultUser               string `xml:"default_user"`
  DefaultGroup              string `xml:"default_group"`
  IntegrationMode           string `xml:"integration_mode"`
  RubyBinPath               string `xml:"ruby"`
  USTRouterAddress          string `xml:"ust_router_address"`
  USTRouterUsername         string `xml:"ust_router_username"`
  USTRouterPassword         string `xml:"ust_router_password"`
  Debugger                  string `xml:"debugger"`
  Analytics                 string `xml:"analytics"`
  MinProcesses              string `xml:"min_processes"`
  MaxProcesses              string `xml:"max_processes"`
  MaxPreloaderIdleTime      string `xml:"max_preloader_idle_time"`
  MaxOutOfBandWorkInstances string `xml:"max_out_of_band_work_instances"`
}

type Process struct {
  PID                 string `xml:"pid"`
  StickySessionID     string `xml:"sticky_session_id"`
  GUPID               string `xml:"gupid"`
  Concurrency         string `xml:"concurrency"`
  Sessions            string `xml:"sessions"`
  Busyness            string `xml:"busyness"`
  RequestsProcessed   string `xml:"processed"`
  SpawnerCreationTime string `xml:"spawner_creation_time"`
  SpawnStartTime      string `xml:"spawn_start_time"`
  SpawnEndTime        string `xml:"spawn_end_time"`
  LastUsed            string `xml:"last_used"`
  LastUsedDesc        string `xml:"last_used_desc"`
  Uptime              string `xml:"uptime"`
  CodeRevision        string `xml:"code_revision"`
  LifeStatus          string `xml:"life_status"`
  Enabled             string `xml:"enabled"`
  HasMetrics          string `xml:"has_metrics"`
  CPU                 string `xml:"cpu"`
  RSS                 string `xml:"rss"`
  PSS                 string `xml:"pss"`
  PrivateDirty        string `xml:"private_dirty"`
  Swap                string `xml:"swap"`
  RealMemory          string `xml:"real_memory"`
  VMSize              string `xml:"vmsize"`
  ProcessGroupID      string `xml:"process_group_id"`
  Command             string `xml:"command"`
}
