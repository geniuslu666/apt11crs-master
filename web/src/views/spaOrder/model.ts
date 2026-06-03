import {h, ref} from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import {isNullObject} from "@/utils/is";
import {NTag, NTooltip} from "naive-ui";

export class State {
  public id = 0; // id
  public ispId = null;
  public orderSn = ''; // 订单编号
  public outOrderSn = ''; // 三方订单号
  public serviceType = 1; // 1到店  2上门
  public confirmType = 1; // 1手动确认  2自动确认
  public memberId = 0; // 用户ID
  public memberDetail = {
    fullName: '',
    memberNo: ''
  };
  public bookDate = ''; // 预定日期
  public bookTime = ''; // 预定时间
  public bookStartTime = null; // 预定开始日期时间
  public bookEndTime = null; // 预定结束日期时间
  public actualStartTime = null; // 实际开始日期时间
  public actualEndTime = null; // 实际结束日期时间
  public bookingName = ''; // 预定人姓名
  public phoneArea = ''; // 区号
  public bookingMobile = ''; // 预定人手机
  public bookingEmail = ''; // 预定人邮箱
  public serviceId = 0; // 服务ID
  public serviceDetail = {
    name: '',
  };
  public orderAmount = 0; // 订单金额
  public couponAmount = 0; // 优惠券抵扣金额
  public balAmount = 0; // 积分抵扣金额
  public goodsId = 0; // 项目ID
  public goodsDetail = {
    goodsName: '',
  };
  public adminMember = {
    username: ''
  };
  public goodsNum = 0; // 人数
  public technicianIds = ''; // 技师IDs
  public technicianList = [
    {
      technicianDetail: {
        nickname: ''
      }
    }
  ];
  public payModel = 3; // 1、余额支付 2、组合支付 3、纯外部支付
  public payTime = null; // 支付时间
  public technicianGoTime = null;
  public memberArriveTime = null; // 预约人到店时间
  public payStatus = 'WAIT'; // WAIT 待支付   HAVE_PAID   已支付    CANCEL  已取消
  public storeId = 0; // 门店ID
  public storeDetail = {
    name: '',
    detailAddress: ''
  };
  public propertyId = 0; // 物业ID
  public propertyDetail = {
    name: '',
  };
  public roomNo = ''; // 房间号
  public orderStatus = 'WAIT_PAY'; // 订单状态  WAIT_PAY 待支付   WAIT_CONFIRM 待确认   WAIT_SERVE 待服务   SERVING 服务中  DONE 已完成   CANCEL 已取消
  public confirmTime = null; // 确认时间
  public confirmRefuseReason = ''; // 审核拒绝原因
  public dispatchStatus = 'WAIT'; // 调度状态 WAIT DONE
  public dispatchTime = null; // 调度时间
  public dispatchDesc = ''; // 调度备注
  public dispatchOperatorId = 0; // 调度员ID
  public expirationTime = 0; // 支付过期时间
  public memberMessage = ''; // 购买人留言信息
  public memberMessageJa = ''; // 购买人留言信息日语
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
  public spaGetRateVip = 0; // 按摩结算积分比例
  public spaGetRateScene = 0; // 场景结算积分比例
  public spaGetScoreStatus = 'WAIT'; // 'WAIT','SUCCESS','FAIL'
  public spaGetAmount = 0;// 结算积分金额
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间
  public adminCancelNum = -1; // 后台取消次数
  public abnormalStatus = 1;
  public abnormalReason = '';
  public memberDeleted = false;

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
    field: 'technicianName',
    component: 'NInput',
    label: '服务技师',
    componentProps: {
      placeholder: '请输入技师名称',
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
]);

// 表格列
export const columns = [
  {
    title: '预约单号',
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
            null,
            {
              default: () => row.orderSn
            }
          ),
          /*h(
            'div',
            null,
            {
              default: () => row.ispId > 0 ? row.spaIspDetail.name : '--',
            }
          ),*/
          h(
            'div',
            {
              style:{
                display: 'flex',
                alignItem: 'center'
              }
            },
            [
              h(
                'div',
                null,
                {
                  default: () => row.ispId > 0 ? row.spaIspDetail.name : '--'
                }
              ),
              htmlStr
            ]
          )
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
    title: '预约人信息',
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
            default: () => '会员已注销'
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
    title: '预约时间',
    key: 'bookStartTime',
    align: 'left',
    width: 180,
  },
  {
    title: '服务地址',
    key: 'serviceType',
    align: 'left',
    width: 180,
    render(row){
      if(row.serviceType == 1){
        return '到店';
      }else{
        return '上门：' + row.propertyDetail.name
      }
    }
  },
  {
    title: '预约项目信息',
    key: 'service',
    align: 'left',
    width: 180,
    ellipsis: false,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => row.serviceDetail.name,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => row.goodsDetail.goodsName + '*' + row.goodsNum + '次',
            }
          )
        ]
      )
    }
  },
  {
    title: '技师信息',
    key: 'technician',
    align: 'left',
    width: 150,
    render(row){
      if(row.dispatchStatus == 'DONE'){
        let htmlArr: any[] = [];
        row.technicianList.forEach((item)=>{
          let htmlItem = h(
            'div',
            null,
            {
              default: () => item.technicianDetail.nickname
            }
          )
          let htmlItem1 = h(
            'div',
            null,
            {
              default: () => item.technicianDetail.phoneArea + item.technicianDetail.phone
            }
          )
          htmlArr.push(htmlItem,htmlItem1)
        })
        return h(
          'div',
          null,
          htmlArr
        )
      }else{
        return '待指定';
      }
    }
  },
  {
    title: '订单总费用',
    key: 'orderAmount',
    align: 'left',
    width: 100,
    render(row){
      return row.orderAmount+'JPY';
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
];

// 字典数据选项
export const options = ref({
  spa_order_status: [] as Option[],
  spa_order_pay_status: [] as Option[],
  work_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['spa_order_status','spa_order_pay_status','work_status'],
  }).then((res) => {
    options.value = res;
  });
}


