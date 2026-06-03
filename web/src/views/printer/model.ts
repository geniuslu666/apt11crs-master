import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import {FormSchema} from "@/components/Form";
import {NTag} from 'naive-ui';

export class State {
  public id = 0; // id
  public printerName = ''; // 打印机名称
  public printerType = 'XPYUN503'; // 打印机类型
  public clientId = null; // 易联云第三方应用ID
  public clientSecret = ''; // 易联云第三方应用秘钥
  public machineCode = ''; // 打印机终端号
  public machineKey = ''; // // 打印机终端密钥
  public printTimes = 1; // 打印联数(次数)
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
  seatNum: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入座位数',
  },
  passengerNum: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入建议乘坐人数',
  },
  maxPackageNum: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入最大容纳行李件数',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'printerName',
    component: 'NInput',
    label: '打印机名称',
    componentProps: {
      placeholder: '请输入打印机名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'machineCode',
    component: 'NInput',
    label: '打印机终端号',
    componentProps: {
      placeholder: '请输入打印机名终端号',
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
    title: '打印机名称',
    key: 'printerName',
    align: 'left',
    width: 130,
  },
  {
    title: '打印机类型',
    key: 'printerType',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.printerType)) {
        return ``;
      }
      if(row.printerType == "YILINK"){
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: "success",
            bordered: false,
          },
          {
            default: () => "易联云",
          }
        );
      }else{
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: "info",
            bordered: false,
          },
          {
            default: () => "芯烨云",
          }
        );
      }
    },
  },
  {
    title: '第三方应用',
    key: 'clientId',
    align: 'left',
    width: 200,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => '应用ID：' +row.clientId,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => '应用秘钥：' + row.clientSecret,
            }
          ),
        ]
      )
    }
  },
  {
    title: '终端号',
    key: 'machineCode',
    align: 'left',
    width: 100,
  },
  {
    title: '终端秘钥',
    key: 'machineKey',
    align: 'left',
    width: 100,
  },
  {
    title: '打印联数',
    key: 'printTimes',
    align: 'left',
    width: 80,
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
    width: 80,
  },
  {
    title: '时间',
    key: 'createAt',
    align: 'left',
    width: 195,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => '创建时间：' + row.createAt,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => '更新时间：' + row.updateAt,
            }
          )
        ]
      )
    }
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


