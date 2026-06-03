// 表格列
import {h, ref} from "vue";
import { FormSchema } from '@/components/Form';
import {NTag, NTooltip} from "naive-ui";


// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'orderSn',
    component: 'NInput',
    label: '订单编号',
    componentProps: {
      placeholder: '请输入订单编号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

export const columns = [
  {
    title: '订单信息',
    key: 'orderSn',
    align: 'left',
    width: 200,
    render(row){
      let htmlStr = ''
      if(row.abnormalStatus == 2){
        htmlStr = h(
          NTag,
          {
            style: {
              marginLeft: '6px',
            },
            type: row.abnormalStatus == 2 ? 'error' : 'warning',
            bordered: false,
            size: 'small'
          },
          {
            default: () => row.abnormalStatus == 2 ? '异常待处理' : '异常已处理',
          }
        )
      }else if(row.abnormalStatus == 3){
        htmlStr = h(
          NTooltip,
          null,
          {
            trigger:()=>
              h(
                NTag,
                {
                  style: {
                    marginLeft: '6px',
                  },
                  type: row.abnormalStatus == 2 ? 'error' : 'warning',
                  bordered: false,
                  size: 'small'
                },
                {
                  default: () => row.abnormalStatus == 2 ? '异常待处理' : '异常已处理',
                }
              ),
            default: () => row.abnormalReason,
          },
        )
      }
      return h(
        'div',
        null,
        [
          h(
            'div',
            {
              style:{
                display: 'flex',
                alignItem: 'center'
              }
            },
            [
              h(
                'div',
                null,
                {
                  default: () => row.orderType == 'INNN' ? 'INNN' : '住一自营'
                }
              ),
              htmlStr
            ]
          ),
          h(
            'div',
            null,
            {
              default: () => row.orderSn
            }
          )
        ]
      )
    }
  },
  {
    title: '订单金额',
    key: 'orderAmount',
    align: 'left',
    width: 200,
  },
  {
    title: '结算比例',
    key: 'settlementRate',
    align: 'left',
    width: 200,
  },
  {
    title: '结算金额',
    key: 'settlementAmount',
    align: 'left',
    width: 200,
  },
  {
    title: '实际结束时间',
    key: 'actualEndTime',
    align: 'left',
    width: 180,
  },
  {
    title: '结算时间',
    key: 'settlementTime',
    align: 'left',
    width: 180,
  },
];
