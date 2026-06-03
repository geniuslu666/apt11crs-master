import {h, ref} from 'vue';
import {NTag, NTooltip} from 'naive-ui';
import { FormSchema } from '@/components/Form';
import {getOptionLabel, getOptionTag, Option} from "@/utils/hotgo";
import {Dicts} from "@/api/dict/dict";
import {isNullObject} from "@/utils/is";
import {defRangeShortcuts} from "@/utils/dateUtil";
import defaultImg from "@/assets/images/mrtx.png";

// 表格搜索表单
export const cabinetOrderSchemas = ref<FormSchema[]>([
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
    {
      field: 'mchBranchName',
      component: 'NInput',
      label: '网点名称',
      componentProps: {
        placeholder: '请输入网点名称',
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'cabinetName',
      component: 'NInput',
      label: '储物柜名称',
      componentProps: {
        placeholder: '请输入储物柜名称',
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'createdAt',
      component: 'NDatePicker',
      label: '下单时间',
      componentProps: {
        type: 'datetimerange',
        clearable: true,
        shortcuts: defRangeShortcuts(),
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'orderStatus',
      component: 'NSelect',
      label: '订单状态',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择订单状态',
        options: [],
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'boxNo',
      component: 'NInput',
      label: '格口编号',
      componentProps: {
        placeholder: '请输入格口编号',
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    /*{
      field: 'memberSearch',
      component: 'NInput',
      label: '会员信息',
      componentProps: {
        placeholder: '请输入会员名/手机号/邮箱',
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },*/
]);

// 表格列
export const cabinetOrderColumns = [
  {
      title: 'ID',
      key: 'id',
      align: 'left',
      width: 50,
    },
    {
      title: '订单编号',
      key: 'orderSn',
      align: 'left',
      width: 120,
    },
    {
      title: '租借网点',
      key: 'mchBranchName',
      align: 'left',
      width: 100,
    },
    {
      title: '储物柜',
      key: 'cabinetName',
      ellipsis: true,
      align: 'left',
      width: 120,
    },
    {
      title: '会员信息',
      key: 'memberNo',
      align: 'left',
      width: 190,
      // resizable: true,
      render(row) {
        return h(
          'div',
          {
            class: 'flex-row',
          },
          [
            h(
              'img',
              {
                src: row.memberInfo.avatar ? row.memberInfo.avatar : defaultImg,
                loading: 'lazy',
                style: {
                  width: '45px',
                  height: '45px',
                  borderRadius: '50%'
                },
                onError: (e) => {
                  e.target.src = defaultImg
                }
              }
            ),
            h(
              'div',
              {
                class:'flex-item',
                style: {
                  paddingLeft: '10px',
                  lineHeight: '22px'
                }
              },
              [
                h(
                  'div',
                  [
                    h(
                      'span',
                      {
                        class: 'cblue f14 fw'
                      },
                      {
                        default: () => row.memberInfo.memberNo,
                      }
                    ),
                  ]
                ),
                h(
                  'div',
                  {},
                  {
                    default: () => row.memberInfo.fullName,
                  }
                )
              ]
            )
          ]
        );
      },
    },
    {
      title: '格口信息',
      key: 'boxNo',
      align: 'left',
      width: 100,
      render(row){
        return h(
          'div',
          null,
          [
            h(
              'div',
              null,
              {
                default: () => row.boxTypeName
              }
            ),
            h(
              'div',
              null,
              {
                default: () => row.boxNo ? row.boxNo : '--',
              }
            )
          ]
        )
      }
    },
    {
      title: '使用时长',
      key: 'useTime',
      align: 'left',
      width: 120,
      render(row) {

        if(row.useTime == 0){
          return '--';
        }

        const hours = Math.floor(row.useTime / 3600);
        const minutes = Math.floor((row.useTime % 3600) / 60);
        const seconds = row.useTime % 60;

        return `${hours}时 ${minutes}分 ${seconds}秒`;
      }
    },
    {
      title: '订单状态',
      key: 'orderStatus',
      align: 'left',
      width: 90,
      render(row) {
        if (isNullObject(row.orderStatus)) {
          return ``;
        }
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: getOptionTag(options.value.cabinet_order_status, row.orderStatus),
            bordered: false,
          },
          {
            default: () => getOptionLabel(options.value.cabinet_order_status, row.orderStatus),
          }
        );
      },
    },
    {
      title: '退款状态',
      key: 'refundStatus',
      align: 'left',
      width: 90,
      render(row) {
        if (isNullObject(row.orderStatus)) {
          return ``;
        }
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: getOptionTag(options.value.refund_status, row.refundStatus),
            bordered: false,
          },
          {
            default: () => getOptionLabel(options.value.refund_status, row.refundStatus),
          }
        );
      },
    },
    {
      title: '创建时间',
      key: 'createdAt',
      align: 'left',
      sorter: true, // 单列排序
      width: 180,
    },
];

// 字典数据选项
export const options = ref({
  cabinet_order_status: [] as Option[],
    refund_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['cabinet_order_status', 'refund_status'],
  }).then((res) => {
    options.value = res;
    for (const item of cabinetOrderSchemas.value) {
      switch (item.field) {
        case 'orderStatus':
                  item.componentProps.options = options.value.cabinet_order_status;
                  break;
      }
    }
  });
}

