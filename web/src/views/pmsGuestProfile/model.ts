import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';

export class State {
  public id = 0; // 主键
  public uid = ''; // 客户档案ID
  public firstName = ''; // 名
  public lastName = ''; // 姓
  public firstNameKana = ''; // 名的假名
  public lastNameKana = ''; // 姓的假名
  public fullName = ''; // 全名
  public language = ''; // 语言
  public email = ''; // 电子邮件
  public phone = ''; // 电话
  public nationality = ''; // 国籍
  public address = ''; // 地址
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间

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
export const rules = {
  email: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.email,
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'fullName',
    component: 'NInput',
    label: '全名',
    componentProps: {
      placeholder: '请输入全名',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'email',
    component: 'NInput',
    label: '电子邮件',
    componentProps: {
      placeholder: '请输入电子邮件',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'phone',
    component: 'NInput',
    label: '电话',
    componentProps: {
      placeholder: '请输入电话',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'nationality',
    component: 'NInput',
    label: '国籍',
    componentProps: {
      placeholder: '请输入国籍',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
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
    title: '主键',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '客户档案ID',
    key: 'uid',
    align: 'left',
    width: -1,
  },
  {
    title: '名',
    key: 'firstName',
    align: 'left',
    width: -1,
  },
  {
    title: '姓',
    key: 'lastName',
    align: 'left',
    width: -1,
  },
  {
    title: '名的假名',
    key: 'firstNameKana',
    align: 'left',
    width: -1,
  },
  {
    title: '姓的假名',
    key: 'lastNameKana',
    align: 'left',
    width: -1,
  },
  {
    title: '全名',
    key: 'fullName',
    align: 'left',
    width: -1,
  },
  {
    title: '语言',
    key: 'language',
    align: 'left',
    width: -1,
  },
  {
    title: '电子邮件',
    key: 'email',
    align: 'left',
    width: -1,
  },
  {
    title: '电话',
    key: 'phone',
    align: 'left',
    width: -1,
  },
  {
    title: '国籍',
    key: 'nationality',
    align: 'left',
    width: -1,
  },
  {
    title: '地址',
    key: 'address',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
];