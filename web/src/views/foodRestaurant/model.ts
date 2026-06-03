import { h, ref } from 'vue';
import {NTag, NSwitch, NInput} from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import {Switch, Sort} from "@/api/foodRestaurant";

const $message = window['$message'];

export class State {
  public id = 0; // id
  public name = ''; // 名称
  public nameLanguage = null;
  public cuisineIds = ''; // 菜系（多选）
  public cuisineIdsArr = [];
  public labelIds = ''; // 标签（多选）
  public labelIdsArr = [];
  public cooperateTypeId = null; // 合作类型ID
  public toretaId = ''; // toreta的餐厅id
  public logo = ''; // LOGO
  public images = ''; // 图集
  public imagesArr = [];
  public content = ''; // 简介
  public contentLanguage = null;
  public phone = ''; // 联系电话
  public openTime = ''; // 营业时间
  public areaPid = null; // 地区省级ID
  public areaId = null; // 地区市级ID
  public detailAddress = ''; // 详细地址
  public ggLat = '34.67100087743556'; // 谷歌纬度
  public ggLng = '135.49982492658245'; // 谷歌经度
  public lat = '39.923678'; // 纬度
  public lng = '116.403478'; // 经度
  public openStatus = null; // 营业状态
  public status = 1; // 1、启用 2、禁用
  public orderTimeType = 1; // 预约日期  1每天 2自定义
  public orderTimeWeek = ''; // 预约日期自定义周数
  public restTimeWeek = ''; // 定休日周数
  public restTimeDate = ''; // 定休日固定日期
  public orderConfirmType = 1; // 预约确定模式 1手动确认 2自动确认
  public orderConfirmDayType = 1; // 1 手动确认在前   2自动确认在前
  public orderConfirmBeforeDays = 0; // 预约确认自定义小时数
  public orderConfirmTime = 0; // 预约确认自定义小时数
  public openTimeType = 1; // 1 按时段  2 按时间点
  public amOpenTime = null; // 早市开始时间
  public amCloseTime = null; // 早市结束时间
  public pmOpenTime = null; // 晚市开始时间
  public pmCloseTime = null; // 晚市结束时间
  public timePoints = ''; // 时间点
  public timeDuration = 15; // 时间间隔
  public orderMaxType = 1; // 最大容纳数  1按时段  2按时间点
  public timeDurationMax = 0; // 时段最大容纳预定数量  0不限制
  public advanceOrderDay = 0; // 至少提前预约天数  0当日可约
  public maxOrderDay = 0; // 最长预约天数
  public cancelPolicyOpen = 1; // 是否允许取消  1允许  2不允许
  public beforeConfirmCancelRate = 0; // 订单确认前取消费用为订单的%，0则免费
  public afterConfirmCancelDay = 0; // 订单确认后距离到店多少天
  public afterConfirmCancelRate1 = 0; // 订单确认后距离到店天数前取消费用为订单的%，0则免费
  public afterConfirmCancelRate2 = 0; // 订单确认后距离到店天数后取消费用为订单的%，0则免费
  public maxSeat = 0; // 最大席位数
  public canSmoking = 1; // 是否允许抽烟  1允许  2不允许
  public settlementType = 1; // 门店抽成类型 1跟随系统  2自定义
  public settlementId = null; // 结算模式ID
  public settlementRate = 0; // 门店结算比例
  public canOrder = 1; // 是否开启预定
  public terminalIds = ''; // 绑定终端ID
  public terminalList = [];
  public printTerminalIds = ''; // 绑定打印机终端ID
  public handTerminalIds = ''; // 绑定手持终端ID
  public account = ''; // 账号
  public password = ''; // 密码
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间
  public deletedAt = ''; // deleted_at
  public toretaCancelEnable = 0; // Toreta是否允许取消 1 允许  2 不允许
  public toretaCancelLimitDay = 0; // Toreta允许取消几天前
  public toretaCancelLimitTime = ""; // Toreta允许取消几天前
  public dayTimeLimit = "";

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
    message: '请输入名称',
  },
  cuisineIds: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入菜系（多选）',
  },
  labelIds: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入标签（多选）',
  },
  cooperateTypeId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入合作类型ID ',
  },
  logo: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入LOGO',
  },
  phone: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入联系电话',
  },
  openTime: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入营业时间',
  },
  detailAddress: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入详细地址',
  },
  ggLat: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入谷歌纬度',
  },
  ggLng: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入谷歌经度',
  },
  lat: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入纬度',
  },
  lng: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入经度',
  },
  openStatus: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入营业状态',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入1、启用 2、禁用',
  },
  orderTimeType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入预约日期  1每天 2自定义',
  },
  orderTimeWeek: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入预约日期自定义周数',
  },
  restTimeWeek: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入定休日周数',
  },
  restTimeDate: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入定休日固定日期',
  },
  orderConfirmType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入预约确定模式 1手动确认 2自动确认',
  },
  orderConfirmHour: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入预约确认自定义小时数',
  },
  timeDuration: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入时间间隔',
  },
  timeDurationMax: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入时段最大容纳预定数量  0不限制',
  },
  advanceOrderDay: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入至少提前预约天数  0当日可约',
  },
  maxOrderDay: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入最长预约天数',
  },
  cancelPolicyOpen: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入是否允许取消  1允许  2不允许',
  },
  beforeConfirmCancelRate: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入订单确认前取消费用为订单的%，0则免费',
  },
  afterConfirmCancelDay: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入订单确认后距离到店多少天',
  },
  afterConfirmCancelRate1: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入订单确认后距离到店天数前取消费用为订单的%，0则免费',
  },
  afterConfirmCancelRate2: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入订单确认后距离到店天数后取消费用为订单的%，0则免费',
  },
  maxSeat: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入最大席位数',
  },
  canSmoking: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入是否允许抽烟  1允许  2不允许',
  },
  settlementType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入门店抽成类型 1跟随系统  2自定义',
  },
  settlementId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入结算模式ID',
  },
  settlementRate: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入门店结算比例',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '名称',
    componentProps: {
      placeholder: '请输入名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '餐厅信息',
    key: 'name',
    align: 'left',
    width: 200,
    canResize: true,
    render(row){
      var cuisineList = [];
      var labelList = [];
      if(row.cuisineList){
        row.cuisineList.forEach((item) => {
          var cuisineH = h(
            NTag,
            {
              style: {
                marginRight: '6px',
              },
              type: 'success',
              bordered: false,
              size: 'small'
            },
            {
              default: () => item.cuisineDetail.cuisineName,
            }
          )
          cuisineList.push(cuisineH)
        })
      }
      if(row.labelList){
        row.labelList.forEach((item) => {
          var labelH = h(
            NTag,
            {
              style: {
                marginRight: '6px',
              },
              type: 'primary',
              bordered: false,
              size: 'small'
            },
            {
              default: () => item.labelDetail.name,
            }
          )
          labelList.push(labelH)
        })
      }

      return h(
        'div',
        null,
        [
          h(
            'div',
            {
              style: {
                marginBottom: row.cuisineList ? '6px' : '0px'
              }
            },
            {
              default: () => row.name,
            }
          ),
          h(
            'div',
            {
              style: {
                marginBottom: row.labelList ? '6px' : '0px'
              }
            },
            cuisineList
          ),
          h(
            'div',
            null,
            labelList
          )
        ]
      )
    }
  },
  {
    title: '营业类型',
    key: 'cooperateTypeId',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.cooperateTypeId)) {
        return `--`;
      }
      return row.cooperateTypeDetail.typeName
    },
  },
  {
    title: '地址信息',
    key: 'address',
    align: 'left',
    width: 240,
    ellipsis: false,
    render(row) {
      let full_address = '';
      // if(row.areaPid > 0){
      //   full_address = row.pAreaDetail.areaName + '-'
      // }
      // if(row.areaId > 0){
      //   full_address = full_address + row.areaDetail.areaName + '-'
      // }
      if(row.detailAddress){
        full_address = full_address + row.detailAddress
      }
      return h(
        'div',
        null,
        [
          h(
            'div',
            {},
            {
              default: () => '联系电话：' + row.phone ? row.phone : '--',
            }
          ),
          h(
            'div',
            {},
            {
              default: () => '地址：' + full_address,
            }
          )
        ]
      )
    },
  },
  // {
  //   title: '核销码',
  //   key: 'verifyCode',
  //   align: 'left',
  //   width: 100,
  //   render(row) {
  //     return h(
  //       NTag,
  //       {
  //         style: {
  //           marginRight: '6px',
  //         },
  //         type: 'warning',
  //         bordered: false,
  //         size: 'medium'
  //       },
  //       {
  //         default: () => row.verifyCode,
  //       }
  //     )
  //   }
  // },
  {
    title: '套餐信息',
    key: 'package',
    align: 'left',
    width: 140,
    render(row) {
      let goodsStatusOnNum = 0;
      if(row.goodsList){
        row.goodsList.forEach((item) => {
          if(item.goodsState == 1){
            goodsStatusOnNum++;
          }
        })
      }
      return h(
        'div',
        null,
        [
          h(
            'div',
            {},
            {
              default: () => '总数量：' + (row.goodsList ? row.goodsList.length : 0),
            }
          ),
          h(
            'div',
            {},
            {
              default: () => '上架中数量：' + goodsStatusOnNum,
            }
          )
        ]
      )
    },
  },
  {
    title: '预定信息',
    key: 'order',
    align: 'left',
    width: 110,
    render(row) {
      return h(
        'div',
        null,
        [
          h(
            'div',
            {},
            {
              default: () => '总预订：' + row.payOrderNum,
            }
          ),
          h(
            'div',
            {},
            {
              default: () => '待确认：' + row.waitConfirmOrderNum,
            }
          )
        ]
      )
    },
  },
  {
    title: '结算信息',
    key: 'settlement',
    align: 'left',
    width: 200,
    render(row) {
      let settlementRate = 0;
      if(row.settlementType == 1){
        settlementRate = row.settlementDetail ? row.settlementDetail.rate : 0
      }else{
        settlementRate = row.settlementRate
      }
      return h(
        'div',
        null,
        [
          h(
            'div',
            {},
            {
              default: () => '预定数：'+row.settlementOrderNum+' / 总金额：' + row.settlementOrderAmount,
            }
          ),
          h(
            'div',
            {},
            {
              default: () => '结算比例：'+settlementRate+'%',
            }
          ),
          h(
            'div',
            {},
            {
              default: () => '结算金额：' + row.totalSettlementAmount,
            }
          )
        ]
      )
    },
  },
  {
    title: '餐厅状态',
    key: 'status',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.openStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.open_status, row.openStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.open_status, row.openStatus),
        }
      );
    },
  },
  {
    title: '预定状态',
    key: 'canOrder',
    align: 'left',
    width: 140,
    render(row) {
      return h(NSwitch, {
        value: row.canOrder === 1,
        checked: '开启',
        unchecked: '关闭',
        onUpdateValue: function (e) {
          row.canOrder = e ? 1 : 2;
          Switch({ id: row.id, key: 'switch', canOrder: row.canOrder }).then((_res) => {
            $message.success('操作成功');
          });
        },
      });
    },
  },
  {
    title: '排序',
    key: 'sort',
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
];

// 字典数据选项
export const options = ref({
  sys_normal_disable: [] as Option[],
  open_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable', 'open_status'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'status':
          item.componentProps.options = options.value.sys_normal_disable;
          break;
        case 'openStatus':
          item.componentProps.options = options.value.open_status;
          break;
      }
    }
  });
}


