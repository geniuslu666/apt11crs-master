import {  ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';

export class State {
  public id = 0; // 编号
  public label = ''; // 分类名称
  public key = ''; // 分类键
  public value = ''; // 分类值
  public icon = ''; // 分类图标
  public tag = 'default'; // 标签

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

// 表单验证规则

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '编号',
    componentProps: {
      placeholder: '请输入编号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'label',
    component: 'NInput',
    label: '分类名称',
    componentProps: {
      placeholder: '请输入分类名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'key',
    component: 'NInput',
    label: '分类键',
    componentProps: {
      placeholder: '请输入分类键',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '编号',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '分类名称',
    key: 'label',
    align: 'left',
    width: -1,
  },
  {
    title: '分类键',
    key: 'key',
    align: 'left',
    width: -1,
  },
  {
    title: '分类值',
    key: 'value',
    align: 'left',
    width: -1,
  },
  {
    title: '分类图标',
    key: 'icon',
    align: 'left',
    width: -1,
  },
];
