# Volcano Metrics List
|Metric Name|Metric Type|Description|Labels|
| ---------- | ----------- | ----------- | ----------- |
|e2e_scheduling_latency_milliseconds|Histogram|E2e scheduling latency in milliseconds (scheduling algorithm + binding)||
|e2e_job_scheduling_latency_milliseconds|Histogram|E2e job scheduling latency in milliseconds||
|e2e_job_scheduling_duration|Gauge|E2E job scheduling duration|labels=["job_name", "queue", "job_namespace"]|
|plugin_scheduling_latency_microseconds|Histogram|Plugin scheduling latency in microseconds|labels=["plugin", "OnSession"]|
|action_scheduling_latency_microseconds|Histogram|Action scheduling latency in microseconds|labels=["action"]|
|task_scheduling_latency_milliseconds|Histogram|Task scheduling latency in milliseconds||
|schedule_attempts_total|Counter|Number of attempts to schedule pods, by the result. 'unschedulable' means a pod could not be scheduled, while 'error' means an internal scheduler problem.|labels=["result"]|
|pod_preemption_victims|Gauge|Number of selected preemption victims||
|total_preemption_attempts|Counter|Total preemption attempts in the cluster till now||
|unschedule_task_count|Gauge|Number of tasks could not be scheduled|labels=["job_id"]|
|unschedule_job_count|Gauge|Number of jobs could not be scheduled||
|job_share|Gauge|Share value for job|labels=["job_ns", "job_id"]|
|job_retry_counts|Counter|Number of retry counts for job|labels=["job_id"]|
|namespace_share|Gauge|Share value for namespace|labels=["namespace_name"]|
|namespace_weight|Gauge|Weight value for namespace|labels=["namespace_name"]|
|namespace_weighted_share|Gauge|Weighted share value for namespace|labels=["namespace_name"]|
|queue_allocated_milli_cpu|Gauge|Allocated CPU for queue|labels=["queue_name"]|
|queue_allocated_memory_bytes|Gauge|Allocated memory for queue|labels=["queue_name"]|
|queue_request_milli_cpu|Gauge|Request CPU for queue|labels=["queue_name"]|
|queue_request_memory_bytes|Gauge|Request memory for queue|labels=["queue_name"]|
|queue_deserved_milli_cpu|Gauge|Deserved CPU for queue|labels=["queue_name"]|
|queue_deserved_memory_bytes|Gauge|Deserved memory for queue|labels=["queue_name"]|
|queue_share|Gauge|Share value for queue|labels=["queue_name"]|
|queue_weight|Gauge|Weight value for queue|labels=["queue_name"]|
|queue_overused|Gauge|queue is overused or not|labels=["queue_name"]|
|queue_pod_group_inqueue_count|Gauge|The number of Inqueue PodGroup in queue|labels=["queue_name"]|
|queue_pod_group_pending_count|Gauge|The number of Pending PodGroup in queue|labels=["queue_name"]|
|queue_pod_group_running_count|Gauge|The number of Running PodGroup in queue|labels=["queue_name"]|
|queue_pod_group_unknown_count|Gauge|The number of Unknown PodGroup in queue|labels=["queue_name"]|