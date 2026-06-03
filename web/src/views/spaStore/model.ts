import { ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';

export class State {
  public id = 0; // id
  public name = ''; // 名称
  public headName = ''; // 负责人
  public phoneArea = ''; // 区号
  public phone = ''; // 联系电话
  public openTime = ''; // 营业时间
  public areaPid = 0; // 地区省级ID
  public areaId = 0; // 地区市级ID
  public detailAddress = ''; // 详细地址
  public ggLat = '34.67100087743556'; // 谷歌纬度
  public ggLng = '135.49982492658245'; // 谷歌经度
  public lat = '39.923678'; // 纬度
  public lng = '116.403478'; // 经度
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间
  public deletedAt = ''; // deleted_at

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
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入门店名称',
  },
  headName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入门店负责人',
  },
  phoneArea: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入区号',
  },
  phone: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入门店电话',
  },
  openTime: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入营业时间',
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
]);

// 表格列
export const columns = [
  {
    title: '门店名称',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: '门店负责人',
    key: 'headName',
    align: 'left',
    width: -1,
  },
  {
    title: '门店电话',
    key: 'phone',
    align: 'left',
    width: -1,
  },
  {
    title: '营业时间',
    key: 'openTime',
    align: 'left',
    width: -1,
  },
  {
    title: '门店地址',
    key: 'area',
    align: 'left',
    width: -1,
  },
  {
    title: '地区市级ID',
    key: 'areaId',
    align: 'left',
    width: -1,
  },
  {
    title: '详细地址',
    key: 'detailAddress',
    align: 'left',
    width: -1,
  },
  {
    title: '谷歌纬度',
    key: 'ggLat',
    align: 'left',
    width: -1,
  },
  {
    title: '谷歌经度',
    key: 'ggLng',
    align: 'left',
    width: -1,
  },
  {
    title: '纬度',
    key: 'lat',
    align: 'left',
    width: -1,
  },
  {
    title: '经度',
    key: 'lng',
    align: 'left',
    width: -1,
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


