import { h, ref } from 'vue';
import { NTag, NInput} from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import {getlang} from "@/utils/smjcomm";
import { useUserStore } from '@/store/modules/user';
import {Sort} from "@/api/pmsProperty";

const userStore = useUserStore();

const $message = window['$message'];

export class State {
  public id = 0; // 主键
  public uid = ''; // 物业ID
  public groupIds = ''; // 绑定会员分组
  public icon = ''; // 图标
  public name = ''; // 物业名称   多语言
  public style = ''; // 物业类型
  public cover = ''; // 封面
  public currency = ''; // 货币
  public language = ''; // 语言
  public timeZone = ''; // 时区
  public tagList = ''; // 标签多语言
  public roomDes = ''; // 房间描述
  public surroundings = ''; // 周边环境
  public description = ''; // 物业描述  多语言
  public busStation = ''; // 公交站
  public subway = ''; // 地铁站
  public contactName = ''; // 联系人
  public phone = ''; // 联系方式
  public contactEmail = ''; // 邮箱
  public maxDaysNotice = 0; // 最大预定区间
  public minutesAfterCheckout = 0; // 退房后 分钟
  public minutesBeforeCheckin = 0; // 入住前 分钟
  public bookingLeadTimeLabel = ''; // 预约期限
  public turnoverDays = 1; // 周转天数
  public checkinAt = ''; // 入住时间
  public checkoutAt = ''; // 退房时间
  public cancelPolicy = ''; // 取消政策 多语言
  public galleryImages = ''; // 画廊图片
  public galleryCover = ''; // 画廊封面
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间
  public close = null; // 状态 1-已启用 2-已关闭
  public regionId = null; // 地区

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
  contactName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入联系人',
  },
  phone: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入联系方式',
  },
  contactEmail: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.email,
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'uid',
    component: 'NInput',
    label: '物业ID',
    componentProps: {
      placeholder: '请输入物业ID',
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
]);

// 表格列
export const columns = [
  {
    title: '序号',
    key: 'index',
    width: 70,
    render(_, index) {
      return index + 1
    }
  },
 /* {
    title: '物业ID',
    key: 'uid',
    align: 'left',
    width: -1,
  },
  {
    title: '图标',
    key: 'icon',
    align: 'left',
    width: 12,
    render(row) {
      if (row.icon === '') {
        return ``;
      }
      return h(
        NAvatar,
        {
          size: 'small',
        },
        {
          default: () => getFileExt(row.icon),
        }
      );
    },
  },*/
  {
    title: '物业名称',
    key: 'name',
    align: 'left',
    width: 300,
    render: function (row){
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

      var leaseCloseLabel = '';
      var bookingCloseLabel = '';
      if(row.leaseClose==1){
        leaseCloseLabel = h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'warning',
            bordered: false,
            size: 'small'
          },
          {
            default: () => "短租",
          }
        )
      }

      if(row.bookingClose==1){
         bookingCloseLabel = h(
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
            default: () => "民宿",
          }
        )
      }


      var propertyName = "--"
      if(row.nameLanguage){
        propertyName = getlang(row.nameLanguage, userStore.language).content
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
              default: () => propertyName,
            }
          ),
          bookingCloseLabel,
          leaseCloseLabel,
          labelH
        ]

      )
    }
  },
  {
    title: '房型数',
    key: 'roomTypeNum',
    align: 'left',
    width: 80,
  },
  {
    title: '房间数',
    key: 'roomUnitNum',
    align: 'left',
    width: 80,
  },
  {
    title: '最小预定区间',
    key: 'minDaysNotice',
    align: 'left',
    width: 100,
  },
  // {
  //   title: '地址',
  //   key: 'address',
  //   align: 'left',
  //   width: -1,
  //   render: function(row){
  //     if (row.addressLanguage) {
  //       return getlang(row.addressLanguage, userStore.language).content;
  //     } else {
  //       return '--';
  //     }
  //   }
  // },
  {
    title: '排序',
    key: 'sort',
    sorter: true, // 单列排序
    width: 100,
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
    title: '状态',
    key: 'name',
    align: 'left',
    width: 90,
    render: function(row){
      if (row.close ==1){
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
            default: () => "已启用",
          }
        );
      }else{
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: "warning",
            bordered: false,
          },
          {
            default: () => "已关闭",
          }
        );
      }

    }
  },
  // {
  //   title: '物业类型',
  //   key: 'style',
  //   align: 'left',
  //   width: -1,
  // },
  // {
  //   title: '封面',
  //   key: 'cover',
  //   align: 'left',
  //   width: 12,
  //   render(row) {
  //     if (row.cover === '') {
  //       return ``;
  //     }
  //     return h(
  //       NAvatar,
  //       {
  //         size: 'small',
  //       },
  //       {
  //         default: () => getFileExt(row.cover),
  //       }
  //     );
  //   },
  // },
  /*{
    title: '货币',
    key: 'currency',
    align: 'left',
    width: -1,
  },
  {
    title: '语言',
    key: 'language',
    align: 'left',
    width: -1,
  },
  {
    title: '时区',
    key: 'timeZone',
    align: 'left',
    width: -1,
  },
  {
    title: '标签多语言',
    key: 'tagList',
    align: 'left',
    width: -1,
  },
  {
    title: '房间描述',
    key: 'roomDes',
    align: 'left',
    width: -1,
  },
  {
    title: '周边环境',
    key: 'surroundings',
    align: 'left',
    width: -1,
  },
  {
    title: '物业描述  多语言',
    key: 'description',
    align: 'left',
    width: -1,
  },
  {
    title: '公交站',
    key: 'busStation',
    align: 'left',
    width: -1,
  },
  {
    title: '地铁站',
    key: 'subway',
    align: 'left',
    width: -1,
  },
  {
    title: '联系人',
    key: 'contactName',
    align: 'left',
    width: -1,
  },
  {
    title: '联系方式',
    key: 'phone',
    align: 'left',
    width: -1,
  },
  {
    title: '邮箱',
    key: 'contactEmail',
    align: 'left',
    width: -1,
  },
  {
    title: '最大预定区间',
    key: 'maxDaysNotice',
    align: 'left',
    width: -1,
  },
  {
    title: '退房后 分钟',
    key: 'minutesAfterCheckout',
    align: 'left',
    width: 12,
  },
  {
    title: '入住前 分钟',
    key: 'minutesBeforeCheckin',
    align: 'left',
    width: 12,
  },
  {
    title: '预约期限',
    key: 'bookingLeadTimeLabel',
    align: 'left',
    width: -1,
  },
  {
    title: '周转天数',
    key: 'turnoverDays',
    align: 'left',
    width: -1,
  },
  {
    title: '入住时间',
    key: 'checkinAt',
    align: 'left',
    width: 12,
  },
  {
    title: '退房时间',
    key: 'checkoutAt',
    align: 'left',
    width: 12,
  },
  {
    title: '取消政策 多语言',
    key: 'cancelPolicy',
    align: 'left',
    width: -1,
  },
  {
    title: '画廊封面',
    key: 'galleryCover',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },*/
];
