import {h, ref} from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import {isNullObject} from "@/utils/is";
import {NIcon, NTag, NTooltip} from "naive-ui";
import {defRangeShortcuts} from "@/utils/dateUtil";
import {QuestionCircleOutlined} from "@vicons/antd";

export class State {
  public canRefund = false; // 是否可退款
  public id = 0; // id
  public orderSn = ''; // 订单编号
  public outOrderSn = ''; // 三方订单号
  public serviceType = 'PICKUP'; // 服务类型
  public orderType = 'PICKUP'; // 订单类型 CRS INNN
  public memberId = 0; // 用户ID
  public memberDetail = {
    fullName: '',
    memberNo: ''
  };
  public bookDate = ''; // 预定日期
  public bookTime = ''; // 预定时间
  public bookStartTime = null; // 预定开始日期时间
  public actualStartTime = null; // 实际开始日期时间
  public actualEndTime = null; // 实际结束日期时间
  public bookingName = ''; // 预定人姓名
  public phoneArea = ''; // 区号
  public bookingMobile = ''; // 预定人手机
  public bookingEmail = ''; // 预定人邮箱
  public adultNum = 0;
  public childNum = 0;
  public startServiceAddress = {
    name: '',
    detailAddress: ''
  };
  public endServiceAddress = {
    name: '',
    detailAddress: ''
  };
  public driverDetail = {
    nickname: ''
  };
  public carDetail = {
    licenseNo: ''
  };
  public serviceDetail = {
    carTypeDetail: {
      name: ''
    }
  }
  public startAddressId = 0; // 出发地ID
  public endAddressId = 0; // 目的地ID
  public serviceId = 0; // 服务ID
  public driverId = 0; // 司机ID
  public carId = 0; // 汽车ID
  public orderAmount = 0; // 订单金额
  public couponAmount = 0; // 优惠券抵扣金额
  public balAmount = 0; // 积分抵扣金额
  public payModel = 3; // 1、余额支付 2、组合支付 3、纯外部支付
  public payTime = null; // 支付时间
  public payStatus = 'WAIT'; // WAIT 待支付   HAVE_PAID   已支付    CANCEL  已取消
  public flightNumber = ''; // 航班号
  public orderStatus = 'WAIT_PAY'; // 订单状态  WAIT_PAY 待支付   WAIT_CONFIRM 待确认   WAIT_SERVE 待服务   SERVING 服务中  DONE 已完成   CANCEL 已取消
  public confirmTime = null; // 确认时间
  public confirmRefuseReason = ''; // 审核拒绝原因
  public driverGoTime = '';
  public dispatchStatus = 'WAIT'; // 调度状态 WAIT DONE
  public dispatchTime = null; // 调度时间
  public dispatchDesc = ''; // 调度备注
  public dispatchOperatorId = 0; // 调度员ID
  public expirationTime = 0; // 支付过期时间
  public startServeImages = ''; // 开始服务图集
  public endServeImages = ''; // 服务结束图集
  public memberMessage = ''; // 购买人留言信息
  public memberMessageJa = ''; // 购买人留言信息日语
  public settlementRate = 0; // 结算比例
  public settlementStatus = 'WAIT'; // WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败
  public settlementAmount = 0; // 结算金额
  public settlementTime = null; // 结算时间
  public settlementOrderId = 0; // 结算单ID
  public settlementType = 1; // 结算方式  1无需结算 2按周期自动结算  3手动申请结算
  public settlementCycle = 1; // 结算周期  1每日结算  2每周结算  3每月结算
  public refundStatus = 'WAIT'; // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
  public refundFee = 0; // 手续费
  public refundTime = ''; // 退款时间
  public refundAmount = 0; // 已退款全部金额
  public refundBalAmount = 0; // 已退款积分金额
  public refundCouponAmount = 0; // 已退款优惠券金额
  public referrer = 0; // 推荐人
  public rebateRate = 0; // 分佣比例
  public rebateStatus = 'WAIT'; // WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败
  public rebateAmount = 0; // 分佣结算金额
  public rebateTime = ''; // 分佣结算时间
  public isGetOpen = 'Y'; // 是否开启积分获取
  public isPayOpen = 'Y'; // 是否开启积分抵扣
  public carGetRateVip = 0; // 接送机结算积分比例
  public carGetRateScene = 0; // 场景结算积分比例
  public carGetScoreStatus = 'WAIT'; // 'WAIT','SUCCESS','FAIL'
  public carGetAmount = 0;// 结算积分金额
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间
  public pickUpSign = '';
  public pickUpSignAmount = 0;
  public childSeatAddNum = 0;
  public childSeatAddAmount = 0;
  public innnLuggageNum = 0;
  public innnLuggageTotalAmount = 0;
  public emergencyName = '';
  public emergencyPhoneArea = '';
  public emergencyMobile = '';
  public adminMember = {
    username: ''
  };
  public adminCancelNum = -1; // 后台取消次数
  public logList = [];
  public abnormalStatus = 1;
  public abnormalReason = '';
  public memberDeleted = false;
  public driverDeleted = false;
  public transactionDetail = [];
  public transactionRefundDetail = [];

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

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'orderSn',
    component: 'NInput',
    label: '订单编号',
    componentProps: {
      placeholder: '请输入订单编号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'serviceType',
    component: 'NSelect',
    label: '类型',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择类型',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'orderType',
    component: 'NSelect',
    label: '渠道',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择渠道',
      options: [
        {
          label: '住一自营',
          value: 'CRS'
        },
        {
          label: 'INNN',
          value: 'INNN'
        },
      ],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'memberSearch',
    component: 'NInput',
    label: '客户信息',
    componentProps: {
      placeholder: '请输入预订人姓名/手机号/用户ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'bookStartTime',
    label: '用车时间',
    slot: 'bookStartTimeSlot',
  },
  {
    field: 'driverName',
    component: 'NInput',
    label: '司机',
    componentProps: {
      placeholder: '请输入司机名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    label: '下单时间',
    slot: 'createdAtSlot',
  },
]);

// 表格列
export const columns = [
  {
    title: '订单信息',
    key: 'orderSn',
    align: 'left',
    width: 180,
    render(row){
      let htmlStr = ''
      if(row.abnormalStatus == 2){
        htmlStr = h(
          NTag,
          {
            style: {
              marginLeft: '6px',
            },
            type: row.abnormalStatus == 2 ? 'error' : 'warning',
            bordered: false,
            size: 'small'
          },
          {
            default: () => row.abnormalStatus == 2 ? '异常待处理' : '异常已处理',
          }
        )
      }else if(row.abnormalStatus == 3){
        htmlStr = h(
          NTooltip,
          null,
          {
            trigger:()=>
              h(
                NTag,
                {
                  style: {
                    marginLeft: '6px',
                  },
                  type: row.abnormalStatus == 2 ? 'error' : 'warning',
                  bordered: false,
                  size: 'small'
                },
                {
                  default: () => row.abnormalStatus == 2 ? '异常待处理' : '异常已处理',
                }
              ),
            default: () => row.abnormalReason,
          },
        )
      }
      return h(
        'div',
        null,
        [
          h(
            'div',
            {
              style:{
              }
            },
            [
              h(
                'div',
                null,
                {
                  default: () => row.orderType == 'INNN' ? 'INNN' : '住一自营'
                }
              ),
              htmlStr
            ]
          ),
          h(
            'div',
            null,
            {
              default: () => row.orderSn
            }
          ),
          h(
            'div',
            {
              style:{
                color: 'red'
              }
            },
            {
              default: () => row.orderAmount + 'JPY'
            }
          ),
        ]
      )
    }
  },
  {
    title: '订单状态',
    key: 'orderStatus',
    align: 'left',
    width: 90,
    render(row) {
      if (isNullObject(row.orderStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.spa_order_status, row.orderStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.spa_order_status, row.orderStatus),
        }
      );
    },
  },
  {
    title: '预约类型',
    key: 'serviceType',
    align: 'left',
    width: 100,
    render(row){
      if(row.serviceType == 'PICKUP'){
        return '接机'
      }else{
        if(row.serviceType == 'DELIVERY'){
          return '送机'
        }else{
          return '包车'
        }
      }
    }
  },
  {
    title: '预约信息',
    key: 'orderSn',
    align: 'left',
    width: 150,
    render(row){
      let memberHtmlStr = ''
      if(row.memberDeleted){
        memberHtmlStr = h(
          'div',
          {
            style:{
              color: 'red'
            }
          },
          {
            default: () => '用户已注销'
          }
        )
      }
      let htmlStr = h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => row.bookingName + ' / ' + row.pmsMemberMemberNo
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.phoneArea+row.bookingMobile,
            }
          ),
          memberHtmlStr
        ]
      )

      return h(
        NTooltip,
        null,
        {
          trigger:()=>
            htmlStr,
          default: () => htmlStr,
        },
      )
    }
  },
  {
    title: '用车时间',
    key: 'bookStartTime',
    align: 'left',
    sorter: true, // 单列排序
    width: 155,
    render(row){
      let datetimeArr = row.bookStartTime.split(' ')
      let dateArr = datetimeArr[0].split('-')
      let timeArr = datetimeArr[1].split(':')
      return parseInt(dateArr[0]) + '-' + parseInt(dateArr[1]) + '-' + parseInt(dateArr[2]) + ' '+parseInt(timeArr[0]) + ':' + timeArr[1]
    }
  },
  {
    title: '线路',
    key: 'address',
    align: 'left',
    width: 240,
    ellipsis: {
      tooltip: true
    },
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => '出发：'+row.startServiceAddress.name
            }
          ),
          h(
            'div',
            null,
            {
              default: () => '到达：'+row.endServiceAddress.name
            }
          )
        ]
      )
    }
  },
  {
    title: '司机',
    key: 'driver',
    align: 'left',
    width: 100,
    render(row){
      if(row.orderType == 'INNN'){
        return '--'
      }else{
        if(row.dispatchStatus == 'WAIT'){
          return '待指派'
        }else{
          let driverHtmlStr = ''
          if(row.driverDeleted){
            driverHtmlStr = h(
              'div',
              {
                style:{
                  color: 'red'
                }
              },
              {
                default: () => '司机已删除'
              }
            )
          }
          return h(
            'div',
            null,
            [
              h(
                'div',
                null,
                {
                  default: () => row.driverDetail.nickname
                }
              ),
              driverHtmlStr
            ]
          )
        }
      }
    }
  },
  {
    title: '支付状态',
    key: 'payStatus',
    align: 'left',
    width: 90,
    render(row) {
      if (isNullObject(row.payStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.spa_order_pay_status, row.payStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.spa_order_pay_status, row.payStatus),
        }
      );
    },
  },
  {
    title: '下单时间',
    key: 'createdAt',
    align: 'left',
    sorter: true, // 单列排序
    width: 180,
  },
];

// 字典数据选项
export const options = ref({
  spa_order_status: [] as Option[],
  spa_order_pay_status: [] as Option[],
  car_service_type: [] as Option[],
  driver_work_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['spa_order_status','spa_order_pay_status','car_service_type','driver_work_status'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'serviceType':
          item.componentProps.options = options.value.car_service_type;
          break;
      }
    }
  });
}


