import { h, ref } from 'vue';
import { NTag, NImage } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { Option, getOptionLabel, getOptionTag, errorImg } from '@/utils/hotgo';

export class State {
  public id = 0; // id
  public language = null; // 语言
  public bannerName = ''; // 名称
  public bannerImage = ''; // 轮播图
  public model = null; // 模块
  public chain = null; // 内外联
  public path = ''; // 链接内容
  public bannerStatus = 1; // 状态
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
  bannerName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入名称',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: 'ID',
    componentProps: {
      placeholder: '请输入ID',
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
    field: 'bannerName',
    component: 'NInput',
    label: '名称',
    componentProps: {
      placeholder: '请输入名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  // {
  //   field: 'model',
  //   component: 'NRadioGroup',
  //   label: '模块',
  //   giProps: {
  //     //span: 24,
  //   },
  //   componentProps: {
  //     options: [],
  //     onUpdateChecked: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  // {
  //   field: 'chain',
  //   component: 'NRadioGroup',
  //   label: '内外联',
  //   giProps: {
  //     //span: 24,
  //   },
  //   componentProps: {
  //     options: [],
  //     onUpdateChecked: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  // {
  //   field: 'bannerStatus',
  //   component: 'NSelect',
  //   label: '状态',
  //   defaultValue: null,
  //   componentProps: {
  //     placeholder: '请选择状态',
  //     options: [],
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
]);

// 表格列
export const columns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '名称',
    key: 'bannerName',
    align: 'left',
    width: -1,
  },
  {
    title: '轮播图',
    key: 'bannerImage',
    align: 'left',
    width: -1,
    render(row) {
      return h(NImage, {
        width: 100,
        src: row.bannerImage,
        onError: errorImg,
        style: {
          'max-width': '100%',
          'max-height': '100%',
          'border-radius': '2px',
        },
      });
    },
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
    title: '模块',
    key: 'model',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.model)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.model, row.model),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.model, row.model),
        }
      );
    },
  },
  {
    title: '内外联',
    key: 'chain',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.chain)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.chain, row.chain),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.chain, row.chain),
        }
      );
    },
  },
  {
    title: '状态',
    key: 'bannerStatus',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.bannerStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.bannerStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.bannerStatus),
        }
      );
    },
  },
  {
    title: '创建时间',
    key: 'createAt',
    align: 'left',
    width: -1,
  },
  {
    title: '更新时间',
    key: 'updateAt',
    align: 'left',
    width: -1,
  },
];

// 字典数据选项
export const options = ref({
  sys_normal_disable: [] as Option[],
  language: [] as Option[],
  chain: [] as Option[],
  model: [] as Option[],
  jump_to_link: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable', 'language', 'chain', 'model', 'jump_to_link'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'bannerStatus':
          item.componentProps.options = options.value.sys_normal_disable;
          break;
        case 'language':
          item.componentProps.options = options.value.language;
          break;
        case 'chain':
          item.componentProps.options = options.value.chain;
          break;
        case 'model':
          item.componentProps.options = options.value.model;
          break;
      }
    }
  });
}


