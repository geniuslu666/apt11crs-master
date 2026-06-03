import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

export class State {
  public id = 0; // id
  public language = null; // 语言
  public categoryId = null; // 分类ID
  public pmsHelpcenterCategoryName = ''; // 分类
  public title = ''; // 标题
  public content = ''; // 内容
  public sort = 0; // 排序
  public createdAt = ''; // created_at
  public updatedAt = ''; // updated_at
  public deletedAt = ''; // deleted_at

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export const rules = {
  title: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入标题',
  },
  content: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入内容',
  },
};

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
    field: 'title',
    component: 'NInput',
    label: '标题',
    componentProps: {
      placeholder: '请输入标题',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'pmsHelpcenterCategoryName',
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
    field: 'createdAt',
    label: '创建时间',
    slot: 'createdAtSlot',
  },
]);

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
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
    title: '分类名称',
    key: 'pmsHelpcenterCategoryName',
    align: 'left',
    width: -1,
  },
  {
    title: '标题',
    key: 'title',
    align: 'left',
    width: -1,
  },
  {
    title: '排序',
    key: 'sort',
    width: 80,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
];

// 字典数据选项
export const options = ref({
  language: [] as Option[],
  category: []
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['language'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'language':
          item.componentProps.options = options.value.language;
          break;
      }
    }
  });
}


