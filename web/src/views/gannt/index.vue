<template>
  <div ref="test"></div>
</template>
<script>
  import 'gantt-schedule-timeline-calendar/dist/style.css';
  import GSTC from 'gantt-schedule-timeline-calendar';

  export default {
    data() {
      return {
        ganttSate: null,
        sc: null,

        licenseKey:
          '====BEGIN LICENSE KEY====\nVgtpnyQnObPC/cjwkxUzy4paHRcxja8UfV3DRw75RALcUBgLBkrlhqZ8QH2ZAo5mrI+460Gd4K/eEupyYO8mrzDOVrWmyv2jfS6b3jIACWAOhxabnWBotuigaAdSmlNNq/P58YtO3kpGdWFywotcQ373Y5q7g4z6FegaRn06I5x/XyefiIFRTFH92hiCuEZn3ckW1tQeo/8LB36nP1INnY4Audmmi0PX31HiykwgHCq2hOlconNmdcQthif7pGHuxW4SPNfNcpuJSuEgi5Nhpxx1dzRLfbkzbPvFdTVrPup2XY2emPpmuU6+SOIbis1QZtw1wYJ8CsOcDpS3Mu0hfQ==||U2FsdGVkX1/O534yt3Gtoc2AWDVnDV7r+/lMurHiDUkeD+67eY+r6H78wRSja5EwdM7nc20BtrlzlpkpY4F8wm+kaTq59XPfZDL8XfLp1uY=\nmPqlt0JklufunWIlJ/c3YEW//uWenq9CW2oRNCOy/hQTbWIo1XyrfBwjd7opDo/BVo9Wl7AMnQ3czMyQVaZ5D3BPJS9nJfMvlAMAYNbq0jaE1x/UC72uWecKPQtJSYWv3EhvZNRxRK3OK5ryuOozSebLM95ICFJL1vC5uNa7TIYmbS05ZjGiduEYCsFNwIAic4vA3kwAb+dY3WbugAk7/6znUSmyoojLql+tS5mKf/ss64phspA/JNhnG12cw8LUdQ6hY6vmRzOTmSymFmUuPUBRzyhPz0p+jRMfYPFTxbUKNksF/zWkbBamwBh6PCOIfG0F0IxyT41Y0LXk2X6OIg==\n====END LICENSE KEY====',
        rowsFromDB: [
          {
            id: 1,
            label: 'Row 1',
          },
          {
            id: 2,
            label: 'Row 2',
          },
        ],
        itemsFromDB: [
          {
            id: '1',
            label: 'Item 1',
            rowId: '1',
            time: {
              start: GSTC.api.date('2020-01-01').startOf('day').valueOf(),
              end: GSTC.api.date('2020-01-02').endOf('day').valueOf(),
            },
          },
          {
            id: '2',
            label: 'Item 2',
            rowId: '1',
            time: {
              start: GSTC.api.date('2020-02-01').startOf('day').valueOf(),
              end: GSTC.api.date('2020-02-02').endOf('day').valueOf(),
            },
          },
          {
            id: '3',
            label: 'Item 3',
            rowId: '2',
            time: {
              start: GSTC.api.date('2020-01-15').startOf('day').valueOf(),
              end: GSTC.api.date('2020-01-20').endOf('day').valueOf(),
            },
          },
        ],
        columnsFromDB: [
          {
            id: 'id',
            data: ({ row }) => GSTC.api.sourceID(row.id), // show original id (not internal GSTCID)
            sortable: ({ row }) => Number(GSTC.api.sourceID(row.id)), // sort by id converted to number
            width: 80,
            header: {
              content: 'ID',
            },
          },
          {
            id: 'label',
            data: 'label',
            sortable: 'label',
            isHTML: false,
            width: 230,
            header: {
              content: 'Label',
            },
          },
        ],
      };
    },
    methods: {
      initGantt() {
        let _this = this;
        const config = {
          licenseKey: _this.licenseKey,
          collapsible: true,
          onExpandCollapse: function (rowId, isExpanded) {
            // 处理展开/折叠事件
            if (isExpanded) {
              console.log(`Row ${rowId} expanded`);
            } else {
              console.log(`Row ${rowId} collapsed`);
            }
          },
          list: {
            columns: {
              data: GSTC.api.fromArray(_this.columnsFromDB),
            },
            rows: GSTC.api.fromArray(_this.rowsFromDB),
          },
          chart: {
            items: GSTC.api.fromArray(_this.itemsFromDB),
          },
        };
        // 从配置对象生成GSTC状态
        this.ganttSate = GSTC.api.stateFromConfig(config);
        // 安装组件
        const app = GSTC({
          element: this.$refs.test,
          state: this.ganttSate,
        });
      },
    },
    mounted() {
      this.initGantt();
      // this.initSocket();
    },
  };
</script>
