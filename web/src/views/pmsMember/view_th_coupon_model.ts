
// 表格列
import {h,ref} from "vue";
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

// 表格搜索表单
export const thCouponSchemas = ref<FormSchema[]>([
  {
    field: 'couponNo',
    component: 'NInput',
    label: '礼品券编号',
    componentProps: {
      placeholder: '请输入编号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'couponName',
    component: 'NInput',
    label: '券名称',
    componentProps: {
      placeholder: '请输入券名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'state',
    component: 'NSelect',
    label: '核销状态',
    componentProps: {
      placeholder: '请选择',
      options: [
        {
          labelField: '未生效',
          valueField: 1,
        },
        {
          labelField: '未使用',
          valueField: 2,
        },
        {
          labelField: '已核销',
          valueField: 3,
        },
        {
          labelField: '已过期',
          valueField: 4,
        },
        {
          labelField: '已失效',
          valueField: 5,
        },
        {
          labelField: '已回收',
          valueField: 6,
        },
      ],
      labelField: 'labelField',
      valueField: 'valueField',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'verifyTime',
    component: 'NDatePicker',
    label: '核销时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

export const thCouponColumns = [
  {
    title: '礼品券编号',
    key: 'couponNo',
    align: 'left',
    width: 155,
  },
  {
    title: '券名称',
    key: 'couponName',
    align: 'left',
    width: 155,
    render(row){
      return row.thCoupon?.couponName
    }
  },
  {
    title: '发放时间',
    key: 'createAt',
    align: 'left',
    width: 150,
    render(row){
      return row.createAt
    }
  },
  {
    title: '生效时间',
    key: 'startTime',
    align: 'left',
    width: 165,
    render(row){
      return h(
        'div',
        [
          h(
            'div',
            [
              h(
                'span',
                {
                  class: 'c999 mr-1'
                },
                {
                  default: () => '开始:',
                }
              ),
              h(
                'span',
                {},
                {
                  default: () => row.startTime ? row.startTime : '--',
                }
              ),
            ]
          ),
          h(
            'div',
            [
              h(
                'span',
                {
                  class: 'c999 mr-1'
                },
                {
                  default: () => '结束:',
                }
              ),
              row.endTime ? row.endTime : '--',
            ]
          )
        ]
      )
    }
  },
  {
    title: '核销状态',
    key: 'state',
    align: 'left',
    width: 150,
    ellipsis: false,
    render(row) {
      if(row.state == 1){
        return '未生效'
      }else if(row.state == 2){
        return '未使用'
      }else if(row.state == 3){
        return '已核销'
      }else if(row.state == 4){
        return '已过期'
      }else if(row.state == 5){
        return h(
          'div',
          {
            class:'flex-item',
            style: {
              // paddingLeft: '8px',
            }
          },
          [
            h(
              'div',
              {
                style: {
                  fontWeight: '400',
                  fontSize: '14px',
                  // color: '#1664FF',
                  lineHeight: '20px'
                }
              },
              {
                default: () => "已失效",
              }
            ),
            h(
              'div',
              {
                style: {
                  fontWeight: '400',
                  fontSize: '14px',
                  color: '#3D3D3D',
                  lineHeight: '20px'
                }
              },
              {
                default: () => row.invalidTime,
              }
            ),
          ]
        )
      }else if(row.state == 6){
        return h(
          'div',
          {
            class:'flex-item',
            style: {
              // paddingLeft: '8px',
            }
          },
          [
            h(
              'div',
              {
                style: {
                  fontWeight: '400',
                  fontSize: '14px',
                  // color: '#1664FF',
                  lineHeight: '20px'
                }
              },
              {
                default: () => "已回收",
              }
            ),
            h(
              'div',
              {
                style: {
                  fontWeight: '400',
                  fontSize: '14px',
                  color: '#3D3D3D',
                  lineHeight: '20px'
                }
              },
              {
                default: () => row.recoveryTime,
              }
            ),
          ]
        )
      }
    }
  },
  {
    title: '核销时间',
    key: 'verifyTime',
    align: 'left',
    width: 155,
  },
  {
    title: '核销商户',
    key: 'verifyMchId',
    align: 'left',
    width: 155,
    render(row){
      return row.verifyMch?.name
    }
  },
  {
    title: '核销门店',
    key: 'verifyStoreId',
    align: 'left',
    width: 155,
    render(row){
      return row.verifyStore?.storeName
    }
  },
  {
    title: '核销商品名',
    key: 'thCouponMch',
    align: 'left',
    width: 155,
    render(row){
      return row.thCouponMchName
    }
  },
];
