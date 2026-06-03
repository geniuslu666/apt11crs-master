import { h, ref } from 'vue';
import {NImage, NTag} from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import {Option, getOptionLabel, getOptionTag, errorImg} from '@/utils/hotgo';

export class State {
  public id = 0; // id
  public language = null; // 语言
  public image = ''; // 主图
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

// 表格搜索表单

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '语言',
    key: 'language',
    align: 'left',
    width: 80,
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
    title: '主图',
    key: 'image',
    align: 'left',
    width: 100,
    render(row) {
      return h(NImage, {
        width: 32,
        height: 32,
        src: row.image,
        onError: errorImg,
        style: {
          width: '32px',
          height: '32px',
          'max-width': '100%',
          'max-height': '100%',
        },
      });
    },
  },
  {
    title: '创建时间',
    key: 'createAt',
    align: 'left',
    width: 180,
  },
  {
    title: '更新时间',
    key: 'updateAt',
    align: 'left',
    width: 180,
  },
];

// 字典数据选项
export const options = ref({
  language: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['language'],
  }).then((res) => {
    options.value = res;
  });
}


