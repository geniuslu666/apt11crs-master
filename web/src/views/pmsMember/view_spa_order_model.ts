import {h, ref} from 'vue';
import {NTag} from 'naive-ui';
import { FormSchema } from '@/components/Form';
import {getOptionLabel, getOptionTag, Option} from "@/utils/hotgo";
import {Dicts} from "@/api/dict/dict";
import {isNullObject} from "@/utils/is";

// 表格搜索表单
export const spaOrderSchemas = ref<FormSchema[]>([
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
    field: 'technicianName',
    component: 'NInput',
    label: '服务技师',
    componentProps: {
      placeholder: '请输入技师名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const spaOrderColumns = [
  {
    title: '预约单号',
    key: 'orderSn',
    align: 'left',
    width: 180,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => row.orderSn
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.ispId > 0 ? row.spaIspDetail.name : '--',
            }
          )
        ]
      )
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
          type: getOptionTag(optionsstatus.value.spa_order_status, row.orderStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(optionsstatus.value.spa_order_status, row.orderStatus),
        }
      );
    },
  },
  {
    title: '预约人信息',
    key: 'orderSn',
    align: 'left',
    width: 150,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => row.bookingName + '/' + row.pmsMemberMemberNo
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.phoneArea+row.bookingMobile,
            }
          )
        ]
      )
    }
  },
  {
    title: '预约时间',
    key: 'bookStartTime',
    align: 'left',
    width: 180,
  },
  {
    title: '服务地址',
    key: 'serviceType',
    align: 'left',
    width: 180,
    render(row){
      if(row.serviceType == 1){
        return '到店';
      }else{
        return '上门：' + row.propertyDetail.name
      }
    }
  },
  {
    title: '预约项目信息',
    key: 'service',
    align: 'left',
    width: 180,
    ellipsis: false,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => row.serviceDetail.name,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.goodsDetail.goodsName + '*' + row.goodsNum + '次',
            }
          )
        ]
      )
    }
  },
  {
    title: '技师信息',
    key: 'technician',
    align: 'left',
    width: 150,
    render(row){
      if(row.dispatchStatus == 'DONE'){
        let htmlArr: any[] = [];
        row.technicianList.forEach((item)=>{
          let htmlItem = h(
            'div',
            null,
            {
              default: () => item.technicianDetail.nickname
            }
          )
          let htmlItem1 = h(
            'div',
            null,
            {
              default: () => item.technicianDetail.phoneArea + item.technicianDetail.phone
            }
          )
          htmlArr.push(htmlItem,htmlItem1)
        })
        return h(
          'div',
          null,
          htmlArr
        )
      }else{
        return '待指定';
      }
    }
  },
  {
    title: '订单总费用',
    key: 'orderAmount',
    align: 'left',
    width: 100,
    render(row){
      return row.orderAmount+'JPY';
    }
  },
  {
    title: '支付状态',
    key: 'payStatus',
    align: 'left',
    width: 90,
    render(row) {
      if (isNullObject(row.payStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(optionsstatus.value.spa_order_pay_status, row.payStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(optionsstatus.value.spa_order_pay_status, row.payStatus),
        }
      );
    },
  },
];

// 字典数据选项
export const optionsstatus = ref({
  spa_order_status: [] as Option[],
  spa_order_pay_status: [] as Option[],
  work_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['spa_order_status','spa_order_pay_status','work_status'],
  }).then((res) => {
    optionsstatus.value = res;
  });
}

