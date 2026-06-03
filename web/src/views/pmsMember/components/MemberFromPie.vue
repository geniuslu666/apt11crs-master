<template>
  <div ref="chartRef" style="width: 100%;height: 300px"></div>
</template>
<script lang="ts">
import {defineComponent, onMounted, ref, Ref, watchEffect} from 'vue';
  import { useECharts } from '@/hooks/web/useECharts';
  import {Stat} from "@/api/pmsMember";

  export default defineComponent({
    props: {
      percentList: {
        type: Array,
        default: []
      }
    },
    setup(props) {
      const chartRef = ref<HTMLDivElement | null>(null);
      const { setOptions } = useECharts(chartRef as Ref<HTMLDivElement>);
      const percentList = ref(props.percentList)
      watchEffect(() => {
        percentList.value = props.percentList;
        setOptions({
          series: [
            {
              type: 'pie',
              radius: ['50%', '80%'],
              avoidLabelOverlap: false,
              itemStyle: {
                borderRadius: 10
              },
              label: {
                show: true,
                formatter: '{b}\n{d}%',
                position: 'outside',
                color: '#999'
              },
              emphasis: {
                label: {
                  show: true,
                  fontSize: 14,
                  fontWeight: 'bold'
                }
              },
              labelLine: {
                show: true,
                length: 15,
                length2: 25,
                smooth: true
              },
              data: percentList.value,
              color: ['#4B87F3','#8BD8FC','#92F2B4','#CEF420'],
            }
          ]
        });
      })

      return { chartRef };
    },
  });
</script>
