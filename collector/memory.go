package collector

import(
    "github.com/prometheus/client_golang/prometheus"
)

//Mem
type MemStats struct {
    Bits                float64 `bson:"bits" type:"counter"`
    Resident            float64 `bson:"resident" type:"counter"`
    Virtual             float64 `bson:"virtual" type:"gauge"`
    Mapped              float64 `bson:"mapped" type:"gauge"`
    MappedWithJournal   float64 `bson:"mappedWithJournal" type:"counter"`
}

func (memStats *MemStats) Collect(exporter *MongodbCollector, ch chan<- prometheus.Metric) {
    group := exporter.FindOrCreateGroupByName("memory")
    group.Collect(memStats, "Bits", ch)
    group.Collect(memStats, "Resident", ch)
    group.Collect(memStats, "Virtual", ch)
    group.Collect(memStats, "Mapped", ch)
    group.Collect(memStats, "MappedWithJournal", ch)
}

