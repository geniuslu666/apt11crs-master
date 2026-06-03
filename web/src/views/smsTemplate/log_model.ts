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
  public templateId = 0; // 系统模版表中主键ID
  public sendTemplateId = ''; // 发送的模版ID
  public type = 1; // 1短信   2邮件   3推送
  public to = ''; // 短信或邮件或推送头
  public vipId = 0; // 会员ID
  public content = ''; // 内容
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
    field: 'type',
    component: 'NSelect',
    label: '类型',
    componentProps: {
      placeholder: '请输入场景',
      options: [
        {
          labelField: '短信',
          valueField: 1,
        },
        {
          labelField: '邮件',
          valueField: 2,
        },
        {
          labelField: '推送',
          valueField: 3,
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
    title: '模版ID',
    key: 'sendTemplateId',
    align: 'left',
    width: 100,
  },
  {
    title: '类型',
    key: 'type',
    align: 'left',
    width: 80,
    render(row){
      if(row.type == 1){
        return '短信'
      }else if(row.type == 2){
        return '邮件'
      }else if(row.type == 3){
        return '推送'
      }
    }
  },
  {
    title: '接收方',
    key: 'to',
    align: 'left',
    width: 140,
  },
  {
    title: '发送内容',
    key: 'content',
    align: 'left',
    width: 200,
    ellipsis: false,
  },
  {
    title: '发送时间',
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


