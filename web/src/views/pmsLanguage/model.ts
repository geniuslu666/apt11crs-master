import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

export class State {
  public id = 0; // id
  public uuid = ''; // 标签ID
  public tag = ''; // 语言标签
  public type = null; // 类型
  public key = ''; // 标识
  public language = null; // 语言
  public content = ''; // 内容
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间

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
  tag: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入语言标签',
  },
  type: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入类型',
  },
  key: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入标识',
  },
  language: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入语言',
  },
  content: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入内容',
  },
};

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
    field: 'uuid',
    component: 'NInput',
    label: '标签ID',
    componentProps: {
      placeholder: '请输入标签ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'tag',
    component: 'NInput',
    label: '语言标签',
    componentProps: {
      placeholder: '请输入语言标签',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'type',
    component: 'NSelect',
    label: '类型',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择类型',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'language',
    component: 'NSelect',
    label: '语言',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择语言',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'content',
    component: 'NInput',
    label: '内容',
    componentProps: {
      placeholder: '请输入内容',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: 100,
  },
  {
    title: '语言标签',
    key: 'tag',
    align: 'left',
    width: -1,
  },
  {
    title: '类型',
    key: 'type',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.type)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.type, row.type),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.type, row.type),
        }
      );
    },
  },
  {
    title: '标识',
    key: 'key',
    align: 'left',
    width: -1,
  },
  {
    title: '语言',
    key: 'language',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.language)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.language, row.language),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.language, row.language),
        }
      );
    },
  },
  {
    title: '内容',
    key: 'content',
    align: 'left',
    width: -1,
  },
];

// 字典数据选项
export const options = ref({
  type: [] as Option[],
  language: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['type', 'language'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'type':
          item.componentProps.options = options.value.type;
          break;
        case 'language':
          item.componentProps.options = options.value.language;
          break;
      }
    }
  });
}
