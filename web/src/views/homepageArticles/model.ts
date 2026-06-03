import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

export interface Coupon {
  couponId: number;
  availableQuantity: number;
}

export interface ThCoupon {
  couponId: number;
  availableQuantity: number;
  perDayAvailable: number;
}

export class State {
  public id = 0; // id
  public language = null; // 语言
  public no = ''; // 编号
  public title = ''; // 标题
  public author = ''; // 作者
  public listPic = ''; // 列表图片
  public startTime = null; // 活动开始时间
  public endTime = null; // 活动结束时间
  public status = 1; // 状态：1-未开始 2-进行中 3-已结束
  public minappStatus = 1; // 状态：1-显示 2-不显示
  public content = ''; // 内容
  public views = 0; // 浏览量
  public thumbNum = 0; // 收藏量
  public sort = 0; // 排序
  public recommendLinkType = ''; // 推荐链接类型 coupon-优惠券 thCoupon-礼品券，hotelDetail-民宿详情，foodIndex-餐厅首页，foodDetail-餐厅详情，spaIndex-按摩首页，spaDetail-按摩详情，carIndex-接送机首页
  public couponLimit = 1; // 每人每张券可领取数量
  public createdAt = ''; // created_at
  public updatedAt = ''; // updated_at
  public deletedAt = ''; // deleted_at
  public couponsArr: Coupon[] = []; // 礼品券idsArr
  public thCouponsArr: ThCoupon[] = []; // 礼品券idsArr
  public hotelIds = ''; // 适用民宿ids
  public restaurantIds = ''; // 适用餐厅ids
  public spaServiceIds = ''; // 适用按摩ids
  public foodIndexTitle = ''; // 餐厅首页标题
  public spaIndexTitle = ''; // 按摩首页标题
  public carIndexTitle = ''; // 接送机首页标题
  public outLinkTitle = ''; // 外链标题
  public outLinkContent = ''; // 外链内容
  public buttonTxt = ''; // 按钮文字
  public couponLinkTitle = ''; // 券链接区域标题
  public couponLinkSubTitle = ''; // 券链接区域副标题
  public recommendLinkTitle = ''; // 推荐链接区域标题
  public recommendLinkSubTitle = ''; // 推荐链接区域副标题
  public chain = 'IN'; // 活动链接方式 IN-内部链接 OUT-外部链接
  public path = ''; // 活动外链链接
  public linkOpenType = ''; // 链接打开方式 1-内部webview 2-外部浏览器
  public limitWeek = ''; // 周末领取限制 6-周六不允许,7-周日不允许，多个用逗号分隔

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export const rules = {
  title: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入标题',
  },
  link: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入链接',
  },
  buttontxt: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入外链按钮文案',
  },
  content: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入内容',
  },
};

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
    field: 'title',
    component: 'NInput',
    label: '标题',
    componentProps: {
      placeholder: '请输入标题',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'language',
    component: 'NSelect',
    label: '语言',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择语言',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    label: '创建时间',
    slot: 'createdAtSlot',
  },
]);

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '语言',
    key: 'language',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.language)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.language, row.language),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.language, row.language),
        }
      );
    },
  },
   {
    title: '活动编号',
    key: 'no',
    align: 'left',
    width: 80,
  },
  {
    title: '标题',
    key: 'title',
    align: 'left',
    width: 180,
  },
 /* {
    title: '作者',
    key: 'author',
    align: 'left',
    width: 120,
  },*/
  {
    title: '浏览量',
    key: 'views',
    align: 'left',
    width: 90,
  },
  {
    title: '点赞量',
    key: 'thumbNum',
    align: 'left',
    width: 90,
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
    width: 80,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 155,
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    align: 'left',
    width: 155,
  },
];

// 字典数据选项
export const options = ref({
  language: [] as Option[],
  sys_normal_disable: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['language', 'sys_normal_disable'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'language':
          item.componentProps.options = options.value.language;
          break;
        case 'status':
          item.componentProps.options = options.value.sys_normal_disable;
          break;
      }
    }
  });
}


