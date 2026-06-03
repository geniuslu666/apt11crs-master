<template>
  <div ref="chartRef" style="width: 100%;height: 294px"></div>
</template>
<script lang="ts">
import {defineComponent, onMounted, ref, Ref, watchEffect} from 'vue';
  import { useECharts } from '@/hooks/web/useECharts';
  import {Stat} from "@/api/pmsMember";
  import * as echarts from "echarts";
  import {hexToRgba} from "@/utils/artDesignUtils";

  export default defineComponent({
    props: {
      dateList: {
        type: Array,
        default: []
      },
      moneyList: {
        type: Array,
        default: []
      },
    },
    setup(props) {
      const chartRef = ref<HTMLDivElement | null>(null);
      const { setOptions } = useECharts(chartRef as Ref<HTMLDivElement>);
      const dateList = ref(props.dateList)
      const moneyList = ref(props.moneyList)
      watchEffect(() => {
        dateList.value = props.dateList;
        moneyList.value = props.moneyList;
        setOptions({
          grid: {
            top: 10,
            right: 20,
            bottom: 10,
            left: 0,
            containLabel: true
          },
          xAxis: {
            type: 'category',
            data: dateList.value,
            boundaryGap: false,
            axisTick: {
              show: false
            },
            axisLine: {
              show: true,
              lineStyle: {
                color: '#e8e8e8',
                width: 1
              }
            },
            axisLabel: {
              show: true,
              color: '#999',
              fontSize: 13
            }
          },
          yAxis: {
            type: 'value',
            axisLabel: {
              show: true,
              color: '#999',
              fontSize: 13,
            },
            axisLine: {
              show: true,
              lineStyle: {
                color: '#e8e8e8',
                width: 1
              }
            },
            splitLine: {
              show: true,
              lineStyle: {
                color: '#e8e8e8',
                width: 1,
                type: 'dashed' as const
              }
            }
          },
          tooltip: {
            trigger: 'axis'
          },
          series: [
            {
              data: moneyList.value,
              type: 'line',
              smooth: true,
              symbol: 'none',
              lineStyle: {
                width: 3,
                color: '#5D87FF'
              },
              areaStyle: {
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                  {
                    offset: 0,
                    color: hexToRgba('#5D87FF', 0.2).rgba
                  },
                  {
                    offset: 1,
                    color: hexToRgba('#5D87FF', 0.01).rgba
                  }
                ])
              },
            }
          ]
        });
      })


      return { chartRef };
    },
  });
</script>
