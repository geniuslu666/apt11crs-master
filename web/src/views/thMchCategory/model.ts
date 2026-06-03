import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import {NSwitch} from "naive-ui";
import {Switch} from "@/api/thMchCategory";

const $message = window['$message'];

export class State {
  public id = 0; // id
  public name = ''; // 分类名称
  public sort = 0; // 排序
  public status = 1; // 状态
  public createdAt = ''; // created_at
  public updatedAt = ''; // updated_at
  public deletedAt = ''; // deleted_at

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: 'id',
    componentProps: {
      placeholder: '请输入id',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: 'created_at',
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
export const columns = [
  {
    title: '分类名称',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
  {
    title: '是否启用',
    key: 'status',
    align: 'left',
    width: -1,
    render(row) {
      return h(NSwitch, {
        value: row.status === 1,
        checked: '开启',
        unchecked: '关闭',
        onUpdateValue: function (e) {
          row.status = e ? 1 : 2;
          Switch({ id: row.id, key: 'switch', status: row.status }).then((_res) => {
            $message.success('操作成功');
          });
        },
      });
    },
  },
];


