// 表格列
import {h, ref} from "vue";
import { FormSchema } from '@/components/Form';


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
    title: '订单编号',
    key: 'orderSn',
    align: 'left',
    width: 200,
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
    title: '订单核销时间',
    key: 'verifyTime',
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
