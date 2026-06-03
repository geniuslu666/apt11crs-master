import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { isNullObject } from '@/utils/is';
import {List as RoomTypeList} from "@/api/pmsRoomType";

export class State {
  public id = 0; // 主键ID
  public planName = ''; //  价格plan名称
  public planShowName = ''; //  价格plan名称
  public planShowNameLanguage = null;
  public propertyId = ''; // 物业UID
  public roomTypeId = ''; // 房型UID
  public bookingDaysType = 'N';
  public bookingDays = null; // 预订天数
  public memberGroupIdType = 'ALL';
  public memberGroupId = null;//用户组
  public memberGroupIdArr = [];
  public memberLevelIdType = 'ALL';
  public memberLevelId = null;//会员等级
  public memberLevelIdArr = [];
  public isCancel = 'Y'; //是否可取消   Y  是  N  否
  public priceMode = '-'; // 模式    +  贵   - 便宜
  public priceStandard = 'PERCENT'; // 基准  PERCENT 倍率  AMOUNT  金额
  public planValue = null; // 价格基准值
  public pricePlanStatus = 'N';//Y 开启  N 关闭
  public createdAt = ''; // create_at
  public updatedAt = ''; // update_at
  public deletedAt = ''; // deleted_at
  public roomTypeDetail = {
    name: ''
  };

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
  planName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入价格计划名称',
  },
  roomTypeId: {
    required: true,
    trigger: 'change',
    message: '请选择房型'
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'planName',
    component: 'NInput',
    label: '价格计划名称',
    componentProps: {
      placeholder: '请输入价格计划名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'roomTypeId',
    component: 'NSelect',
    label: '房型',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择房型',
      options: [],
      labelField: 'name',
      valueField: 'uid',
      onUpdateValue: (e: any) => {
        console.log(typeof(e));
      },
    },
  },
  {
    field: 'pricePlanStatus',
    component: 'NSelect',
    label: '有效状态',
    componentProps: {
      placeholder: '请选择有效状态',
      options: [
        {
          labelField: '有效',
          valueField: 'Y',
        },
        {
          labelField: '已停用',
          valueField: 'N',
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
    title: '价格计划名称',
    key: 'planName',
    align: 'left',
    width: 150,
    render(row) {
      var labelH = ""
      if(row.deletedAt != null){
        labelH = h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'error',
            bordered: false,
            size: 'small'
          },
          {
            default: () => "已删除",
          }
        )
      }
      return h(
        'div',
        null,
        [
          h(
            'span',
            {
              style: {
                display: "inline-block",
                marginRight: '6px'
              }
            },
            {
              default: () => row.planName,
            }
          ),
          labelH
        ]

      )
    }
  },
  {
    title: '取消政策',
    key: 'isCancel',
    align: 'left',
    width: 80,
    render(row){
      if(row.isCancel == "Y"){
        return "灵活取消"
      }else{
        return "不可取消"
      }
    }
  },
  {
    title: '房价Plan',
    key: 'priceMode',
    align: 'left',
    width: 250,
    render(row){
      if(row.priceMode == "+"){
        var priceMode = '贵';
      }else{
        var priceMode = '便宜';
      }
      return '房价比Standard Rate' + priceMode + (row.priceStandard == "AMOUNT" ? 'JPY' : '') + row.planValue + (row.priceStandard == "PERCENT" ? '%' : '')
    }
  },
  {
    title: '关联房型',
    key: 'roomType',
    align: 'left',
    width: 150,
    render(row){
      return row.roomTypeDetail ? row.roomTypeDetail.name : '--'
    }
  },
  {
    title: '有效状态',
    key: 'pricePlanStatus',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.pricePlanStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.pricePlanStatus == 'Y' ? 'success' : 'warning',
          bordered: false,
        },
        {
          default: () => row.pricePlanStatus == 'Y' ? '有效' : '已停用',
        }
      );
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 180,
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    align: 'left',
    width: 180,
  },
];

export const roomTypeList = ref([]);

export function getRoomTypeList(propertyId) {
  RoomTypeList({
    puid: propertyId,
    pagination: false
  }).then((res) => {
    roomTypeList.value = res.list;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'roomTypeId':
          item.componentProps.options = roomTypeList.value;
          break;
      }
    }
  });
}


