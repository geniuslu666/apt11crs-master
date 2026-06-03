import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import {defRangeShortcuts, defShortcuts, formatToDate} from '@/utils/dateUtil';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import defaultImg from "@/assets/images/onerror.png";
import {jsontoobj} from "@/utils/smjcomm";

export class State {
  public id = 0; // id
  public date = ''; // 活动日期
  public name = ''; // 标题（多语言）
  public nameLanguage = null; // 标题-多语言
  public subName = ''; // 副标题
  public pic = ''; // 图片
  public restaurantIds = ''; // 餐厅ids
  public restaurantIdsArr = []; // 餐厅idsArr
  public status = 1; // 状态
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
  subName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入副标题',
  },
  pic: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入图片',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'date',
    component: 'NDatePicker',
    label: '活动日期',
    componentProps: {
      type: 'daterange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
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
    field: 'createAt',
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
    title: 'id',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '活动日期',
    key: 'date',
    align: 'left',
    width: 100,
    render(row) {
      return formatToDate(row.date);
    },
  },
  {
    title: '活动标题',
    key: 'name',
    align: 'left',
    width: 250,
    render(row) {
      let nameLanguage = jsontoobj(row.nameLanguage);
      return nameLanguage.zh && nameLanguage.zh.content ? nameLanguage.zh.content : ''
    },
  },
 /* {
    title: '副标题',
    key: 'subName',
    align: 'left',
    width: -1,
  },*/
  {
    title: '图片',
    key: 'pic',
    align: 'left',
    width: 80,
    render(row) {

      var pic = row.pic;
      if(pic == ""|| !pic){
        // 如果活动图不存在，则获取餐厅第一张图片
        if(!row.restaurantList || row.restaurantList.length <= 0){
           pic = '';
        }else{
          var imageArr = row.restaurantList[0].restaurantDetail.images.split(",").map(item => String(item));
          if(imageArr.length <= 0){
            pic = '';
          }else{
            pic = imageArr[0];
          }
        }
      }

      return h(
        'div',
        {
          class: 'flex-row',
        },
        [
          h(
            'img',
            {
              src: pic ? pic : defaultImg,
              loading: 'lazy',
              style: {
                width: '45px',
                height: '45px',
                // borderRadius: '50%'
              },
              onError: (e) => {
                e.target.src = defaultImg
              }
            }
          ),
        ]
      );
    },
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
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'status':
          item.componentProps.options = options.value.sys_normal_disable;
          break;
      }
    }
  });
}


