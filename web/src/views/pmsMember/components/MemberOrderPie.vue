<template>
  <div ref="chartRef" style="width: 100%;height: 100%"></div>
</template>
<script lang="ts">
import {defineComponent, onMounted, ref, Ref, watchEffect} from 'vue';
  import { useECharts } from '@/hooks/web/useECharts';

  export default defineComponent({
    props: {
      memberOrderPercent: {
        type: Number,
        default: 0
      },
      memberNotOrderPercent: {
        type: Number,
        default: 0
      }
    },
    setup(props) {
      const chartRef = ref<HTMLDivElement | null>(null);
      const { setOptions } = useECharts(chartRef as Ref<HTMLDivElement>);
      const memberOrderPercent = ref(props.memberOrderPercent)
      const memberNotOrderPercent = ref(props.memberNotOrderPercent)

      watchEffect(() => {
        memberOrderPercent.value = props.memberOrderPercent;
        memberNotOrderPercent.value = props.memberNotOrderPercent;

        setOptions({
          series: [
            {
              type: 'pie',
              center: ['50%', '50%'],
              radius: ['40%', '75%'],
              avoidLabelOverlap: false,
              itemStyle: {
                borderRadius: 10
              },
              label: {
                show: false,
                formatter: '{b}\n{d}%',
                position: 'center',
                color: '#999'
              },
              emphasis: {
                label: {
                  show: true,
                  fontSize: 14,
                  fontWeight: 'bold'
                }
              },
              data: [
                { value: memberOrderPercent.value, name: '下单会员数' },
                { value: memberNotOrderPercent.value, name: '未下单会员数' },
              ],
              color: ['#4B87F3','#8BD8FC','#92F2B4','#CEF420'],
            }
          ]
        });
      });

      onMounted(() => {

      });
      return { chartRef };
    },
  });
</script>
