import { h, ref } from 'vue';
import {NTag, NInput, NEllipsis} from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { Dicts } from '@/api/dict/dict';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import { isNullObject } from '@/utils/is';
import {Sort} from "@/api/carService";

const $message = window['$message'];

export class State {
  public id = 0; // id
  public serviceType = ''; // 服务类型
  public serviceName = ''; // 路线名称
  public carTypeId = null; // 车型ID
  public startIds = null; // 出发机场
  public endIds = ''; // 到达物业 多选
  public endIdsType = 1;
  public endIdsArr = [];
  public distance = 0; // 路线长度 KM
  public useTime = 0; // 线路时长  分钟
  public price = 0; // 基础费用
  public linePrice = 0; // 划线价
  public freeWaitTime = 0; // 免费等待时长  分钟
  public maxWaitTime = 0; // 最大等待时长  分钟
  public timeoutPreTime = 0; // 超时每xx分钟
  public timeoutPrePrice = 0; // 超时价格
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
  serviceName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入路线名称',
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
    title: '路线名称',
    key: 'serviceName',
    align: 'left',
    width: 150,
  },
  {
    title: '车型',
    key: 'carTypeId',
    align: 'left',
    width: 100,
    render(row){
      return row.carTypeDetail.name
    }
  },
  {
    title: '出发机场',
    key: 'startIds',
    align: 'left',
    width: 150,
    render(row){
      return row.startServiceAddress[0].addressDetail.name
    }
  },
  {
    title: '目的地',
    key: 'endIds',
    align: 'left',
    width: 150,
    render(row){
      if(!row.endIds){
        return '全部物业'
      }else{
        var htmlArr = [];
        row.endServiceAddress.forEach((item)=>{
          if(item.addressDetail.status == 1){
            var htmlItem = h(
              'div',
              null,
              {
                default: () => item.addressDetail.name
              }
            )
            htmlArr.push(htmlItem)
          }
        })
        return h(
          NEllipsis,
          {
            expandTrigger: 'click',
            lineClamp: 3,
            tooltip: false,
          },
          {
            default: () => htmlArr,
          }
        )
      }
    }
  },
  {
    title: '预估距离',
    key: 'distance',
    align: 'left',
    width: 80,
    render(row) {
      return row.distance + 'KM'
    }
  },
  {
    title: '预估时长',
    key: 'useTime',
    align: 'left',
    width: 80,
    render(row) {
      return row.useTime + '分钟'
    }
  },
  {
    title: '基础费用',
    key: 'price',
    align: 'left',
    width: 90,
    render(row){
      return row.price + 'JPY'
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
    title: '排序',
    key: 'sort',
    align: 'left',
    sorter: true, // 单列排序
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
    width: 160,
  },
  {
    title: '更新时间',
    key: 'updateAt',
    align: 'left',
    width: 160,
  },
];

// 字典数据选项
export const options = ref({
  sys_normal_disable: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable'],
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


