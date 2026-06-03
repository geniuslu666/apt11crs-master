import { h, ref } from 'vue';
import {NImage, NInput, NSwitch, NTag} from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { Dicts } from '@/api/dict/dict';
import {Option, errorImg, getOptionTag, getOptionLabel} from '@/utils/hotgo';
import {Switch, Sort} from "@/api/pmsIndexBanner";
import {isNullObject} from "@/utils/is";

const $message = window['$message'];

export class State {
  public id = 0; // id
  public language = null; // 语言
  public bannerImage = ''; // 轮播图
  public model = null; // 模块
  public chain = "IN"; // 内外联
  public path = ''; // 链接内容
  public bannerStatus = 1; // 状态
  public minappStatus = 1; // 小程序是否显示 1 显示 2 不显示
  public sort = 0; // 排序(越大越靠前)
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间
  public linkOpenType = 1; // 外链跳转方式 1-内部webview 2-外部浏览器

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
    title: '轮播图',
    key: 'bannerImage',
    align: 'left',
    width: 200,
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
  /*{
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
  },*/
  {
    title: '内外联',
    key: 'chain',
    align: 'left',
    width: 100,
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
 /* {
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
  },*/
  {
    title: '是否启用',
    key: 'bannerStatus',
    align: 'left',
    width: 130,
    render(row) {
      return h(NSwitch, {
        value: row.bannerStatus === 1,
        checked: '开启',
        unchecked: '关闭',
        onUpdateValue: function (e) {
          row.bannerStatus = e ? 1 : 2;
          Switch({ id: row.id, key: 'switch', status: row.bannerStatus }).then((_res) => {
            $message.success('操作成功');
          });
        },
      });
    },
  },
  {
    title: '小程序显示',
    key: 'minappStatus',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.minappStatus)) {
        return ``;
      }
      if(row.minappStatus === 1){
        return "显示"
      }else if(row.minappStatus === 2){
        return "不显示"
      }
      return `--`;
    },
  },
  {
    title: '排序',
    key: 'sort',
    width: 80,
    render(row) {
      return h(NInput, {
        value: row.sort,
        onUpdateValue(v) {
          Sort({ id: row.id, sort: v }).then((_res) => {
            $message.success('操作成功');
          });
          row.sort = v
        }
      })
    }
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
  });
}


