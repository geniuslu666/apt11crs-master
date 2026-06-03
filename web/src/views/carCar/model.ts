import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { Dicts } from '@/api/dict/dict';
import { Option,getOptionLabel,getOptionTag } from '@/utils/hotgo';
import { isNullObject } from '@/utils/is';
import {NTag} from 'naive-ui';

export class State {
  public id = 0; // id
  public carName = ''; // 车辆名称
  public typeId = null; // 车型ID
  public brand = ''; // 车辆品牌型号
  public licenseNo = ''; // 车牌号码
  public licenseColor = 'WHITE'; // 牌照颜色
  public qualityMaterials = ''; // 资质材料
  public qualityMaterialsArr = [];
  public status = 1; // 状态
  public sort = 0; // 排序(越大越靠前)
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
  carName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入车辆名称',
  },
  brand: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入车辆品牌型号',
  },
  licenseNo: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入车牌号码',
  },
};

// 表格搜索表单
// export const schemas = ref<FormSchema[]>([
//   {
//     field: 'labelName',
//     component: 'NInput',
//     label: '标签名称',
//     componentProps: {
//       placeholder: '请输入标签名称',
//       onUpdateValue: (e: any) => {
//         console.log(e);
//       },
//     },
//   },
//   {
//     field: 'status',
//     component: 'NSelect',
//     label: '状态',
//     defaultValue: null,
//     componentProps: {
//       placeholder: '请选择状态',
//       options: [],
//       onUpdateValue: (e: any) => {
//         console.log(e);
//       },
//     },
//   },
// ]);

// 表格列
export const columns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '车辆名称',
    key: 'carName',
    align: 'left',
    width: 150,
  },
  {
    title: '车型',
    key: 'desc',
    align: 'left',
    width: 150,
    render(row) {
      return row.carTypeDetail.name
    },
  },
  {
    title: '品牌型号',
    key: 'brand',
    align: 'left',
    width: 150,
  },
  {
    title: '牌照信息',
    key: 'licenseNo',
    align: 'left',
    width: 180,
    render(row){
      let licenseColorName = '未知';
      if(row.licenseColor == "WHITE"){
        licenseColorName = '白牌'
      }else if(row.licenseColor == "GREEN"){
        licenseColorName = '绿牌'
      }else{
        licenseColorName = '黄牌'
      }
      return row.licenseNo + '/' + licenseColorName
    }
  },
  {
    title: '车型图片',
    key: 'image',
    align: 'left',
    width: 100,
    render(row){
      return h(
        'img',
        {
          src: row.carTypeDetail.image,
          loading: 'lazy',
          style: {
            width: '45px',
            height: '45px',
          },
        }
      )
    }
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.status)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.status),
        }
      );
    },
  },
  {
    title: '工作状态',
    key: 'workStatus',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.workStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.car_work_status, row.workStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.car_work_status, row.workStatus),
        }
      );
    },
  },
  // {
  //   title: '排序',
  //   key: 'sort',
  //   align: 'left',
  //   width: 80,
  // },
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
  car_work_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable', 'car_work_status'],
  }).then((res) => {
    options.value = res;
    // for (const item of schemas.value) {
    //   switch (item.field) {
    //     case 'status':
    //       item.componentProps.options = options.value.sys_normal_disable;
    //       break;
    //   }
    // }
  });
}


