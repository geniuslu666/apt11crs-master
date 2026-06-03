import {h, ref} from 'vue';
import { FormSchema } from '@/components/Form';
import {defRangeShortcuts} from '@/utils/dateUtil';
import {GetMemberOption} from "@/api/org/user";

// 表格搜索表单
export const expSchemas = ref<FormSchema[]>([
  {
    field: 'operatorId',
    component: 'NSelect',
    label: '操作人',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择操作人',
      options: [],
      labelField: 'username',
      valueField: 'value',
      filterable: true,
      onUpdateValue: (e: any) => {
        console.log(typeof(e));
      },
    },
  },
  {
    field: 'des',
    component: 'NInput',
    label: '原因',
    componentProps: {
      placeholder: '请输入原因',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '发生时间',
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

// 表格列
export const expColumns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 100,
  },
  {
    title: '方向',
    key: 'exp',
    align: 'left',
    width: 80,
    render(row) {
      return h(
        'span',
        {
          style: {
            color: row.exp > 0 ? 'green' : 'red'
          },
        },
        {
          default: () => row.exp > 0 ? '发放' : '消耗'
        }
      )
    },
  },
  {
    title: '发生场景',
    key: 'scene',
    align: 'left',
    width: 80,
    render(row) {
      if(row.scene == 'SYSTEM'){
        return '系统'
      }else if(row.scene == 'HOTEL'){
        return '酒店'
      }else{
        return '未知'
      }
    },
  },
  {
    title: '成长值变化',
    key: 'exp',
    align: 'left',
    width: 200,
    render(row) {
      return h(
        'span',
        {
          style: {
            color: row.exp > 0 ? 'green' : 'red'
          },
        },
        {
          default: () => row.exp > 0 ? '+' + row.exp : row.exp
        }
      )
    },
  },
  {
    title: '操作人',
    key: 'des',
    align: 'left',
    width: 100,
    render(row){
      if(row.operatorId > 0){
        return row.adminMemberUsername
      }else{
        return '系统操作'
      }
    }
  },
  {
    title: '原因',
    key: 'des',
    align: 'left',
    width: 200,
    ellipsis: false,
  },
  {
    title: '发生时间',
    key: 'createdAt',
    align: 'left',
    width: 200,
  },
  {
    title: '关联订单号',
    key: 'orderSn',
    align: 'left',
    width: 240,
  },
];

export const operatorList = ref([]);

// 加载字典数据选项
export function loadExpOptions() {
  GetMemberOption().then((res) => {
    operatorList.value = res;
    for (const item of expSchemas.value) {
      switch (item.field) {
        case 'operatorId':
          item.componentProps.options = operatorList.value;
          break;
      }
    }
  });
}

