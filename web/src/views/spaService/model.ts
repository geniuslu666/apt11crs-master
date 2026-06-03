import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { Option } from '@/utils/hotgo';

export class State {
  public id = 0; // id
  public ispId = null;// 服务商
  public name = ''; // 服务名称
  public nameLanguage = null;
  public subName = ''; // 简介
  public subNameLanguage = null;
  public labelIds = ''; // 标签（多选）
  public labelIdsArr = [];
  public images = ''; // 图集
  public imagesArr = [];
  public channel = 1; // 渠道 1-到店或上门 2-仅上门 3-仅到店
  public propertyType = 1; // 1-全部物业 2-部分物业
  public propertyIds = ''; // 物业（多选）
  public propertyIdsArr = [];
  public serviceState = 1; // 状态
  public sort = 0; // 排序(越大越靠前)
  public content = ''; // 详情
  public contentLanguage = null;
  public priceList = [];
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间
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

// 表单验证规则
export const rules = {
  sort: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入排序',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '服务名称',
    componentProps: {
      placeholder: '请输入服务名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);


// 字典数据选项
export const options = ref({
  service_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['service_status'],
  }).then((res) => {
    options.value = res;
  });
}


