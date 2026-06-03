import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

export class State {
  public id = 0; // ID
  public scene = 1; // 场景 1会员  2住宿  3接送机  4按摩
  public event = ''; // 别名
  public title = ''; // 标题
  public content = ''; // 内容
  public contentLanguage = null; // 内容-多语言
  public umsId = ''; // 一信通模版ID
  public umsContent = ''; // 一信通模版内容
  public umsParam = ''; // 一信通变量
  public umsTemplate = null; // 一信通JSON
  public tencentJaId = ''; // 腾讯云日文模版ID
  public tencentJaContent = ''; // 腾讯云日文模版内容
  public tencentJaParam = ''; // 腾讯云日文变量
  public tencentKoId = ''; // 腾讯云韩文模版ID
  public tencentKoContent = ''; // 腾讯云韩文模版内容
  public tencentKoParam = ''; // 腾讯云韩文变量
  public tencentEnId = ''; // 腾讯云英文模版ID
  public tencentEnContent = ''; // 腾讯云英文模版内容
  public tencentEnParam = ''; // 腾讯云英文变量
  public tencentTwId = ''; // 腾讯云繁体模版ID
  public tencentTwContent = ''; // 腾讯云繁体模版内容
  public tencentTwParam = ''; // 腾讯云繁体变量
  public tencentTemplate = null; // 一信通JSON
  public aliyunJaId = ''; // 阿里云日文模版ID
  public aliyunJaContent = ''; // 阿里云日文模版内容
  public aliyunJaParam = ''; // 阿里云日文变量
  public aliyunKoId = ''; // 阿里云韩文模版ID
  public aliyunKoContent = ''; // 阿里云韩文模版内容
  public aliyunKoParam = ''; // 阿里云韩文变量
  public aliyunEnId = ''; // 阿里云英文模版ID
  public aliyunEnContent = ''; // 阿里云英文模版内容
  public aliyunEnParam = ''; // 阿里云英文变量
  public aliyunTwId = ''; // 阿里云繁体模版ID
  public aliyunTwContent = ''; // 阿里云繁体模版内容
  public aliyunTwParam = ''; // 阿里云繁体变量
  public aliyunTemplate = null; // 一信通JSON
  public createAt = ''; // 创建时间
  public updateAt = ''; // 修改时间
  public deletedAt = ''; // 删除时间

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
    field: 'scene',
    component: 'NSelect',
    label: '场景',
    componentProps: {
      placeholder: '请输入场景',
      options: [
        {
          labelField: '会员',
          valueField: 1,
        },
        {
          labelField: '住宿',
          valueField: 2,
        },
        {
          labelField: '接送机',
          valueField: 3,
        },
        {
          labelField: '按摩',
          valueField: 4,
        }
      ],
      labelField: 'labelField',
      valueField: 'valueField',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '场景',
    key: 'type',
    align: 'left',
    width: 80,
    render(row){
      if(row.scene == 1){
        return '会员'
      }else if(row.scene == 2){
        return '住宿'
      }else if(row.scene == 3){
        return '接送机'
      }else if(row.scene == 4){
        return '按摩'
      }
    }
  },
  {
    title: '别名',
    key: 'event',
    align: 'left',
    width: 140,
  },
  {
    title: '标题',
    key: 'title',
    align: 'left',
    width: 180,
  },
  {
    title: '通知内容',
    key: 'content',
    align: 'left',
    width: 200,
    ellipsis: false,
  },
  {
    title: '创建时间',
    key: 'createdAt',
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
    for (const item of schemas.value) {
      switch (item.field) {
        case 'language':
          item.componentProps.options = options.value.language;
          break;
      }
    }
  });
}


