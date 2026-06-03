import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import {isNullObject} from "@/utils/is";
import {NTag} from "naive-ui";
import {getOptionLabel, getOptionTag} from "@/utils/hotgo";
import {options} from "@/views/foodActivity/model";

export class State {
  public uid = ''; // 第三方系统的ID
  public id = 0; // 主键
  public rtUid = ''; // 房型ID
  public roomNo = ''; // 房间号
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间

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
  uid: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入第三方系统的ID',
  },
  rtUid: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入房型ID',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'roomNo',
    component: 'NInput',
    label: '房间号',
    componentProps: {
      placeholder: '请输入房间号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'pmsRoomTypeName',
    component: 'NInput',
    label: '房型名称',
    componentProps: {
      placeholder: '请输入房型名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  // {
  //   title: '主键',
  //   key: 'id',
  //   align: 'left',
  //   width: -1,
  // },
  {
    title: '房间号',
    key: 'roomNo',
    align: 'left',
    width: -1,
  },
  {
    title: '状态',
    key: 'pmsRoomTypeIsShow',
    align: 'left',
    width: -1,
    render(row){

      if (isNullObject(row.roomTypeDetail.isShow)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.roomTypeDetail.isShow == 1 ? 'success' : 'error',
          bordered: false,
        },
        {
          default: () =>row.roomTypeDetail.isShow == 1 ? '正常' : '暂停',
        }
      );
    }
  },
  {
    title: '房型名称',
    key: 'pmsRoomTypeName',
    align: 'left',
    width: -1,
    render(row){
      return row.roomTypeDetail.name
    }
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
