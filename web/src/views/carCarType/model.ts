import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { Dicts } from '@/api/dict/dict';
import { Option } from '@/utils/hotgo';

export class State {
  public id = 0; // id
  public name = ''; // 名称
  public nameLanguage = null;
  public desc = ''; // 简介
  public descLanguage = null;
  public image = ''; // 图片
  public seatNum = 0; // 座位数
  public passengerNum = 0; // 建议乘员人数
  public maxPackageNum = 0; // 最大容纳行李件数
  public childrenSeatSupport = 'Y'; // 是否支持儿童座椅
  public childrenSeatSpace = 1; // 儿童座椅占几个作为
  public content = ''; // 创建时间
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
    title: '车型名称',
    key: 'name',
    align: 'left',
    width: 150,
  },
  {
    title: '车型简介',
    key: 'desc',
    align: 'left',
    width: 150,
  },
  {
    title: '车型图片',
    key: 'image',
    align: 'left',
    width: 100,
    render(row){
      return h(
        'img',
        {
          src: row.image,
          loading: 'lazy',
          style: {
            width: '45px',
            height: '45px',
            borderRadius: '50%'
          },
        }
      )
    }
  },
  {
    title: '座位数',
    key: 'seatNum',
    align: 'left',
    width: 80,
    render(row){
      return row.seatNum + '座'
    }
  },
  {
    title: '建议成员人数',
    key: 'passengerNum',
    align: 'left',
    width: 80,
    render(row){
      return row.passengerNum + '人'
    }
  },
  // {
  //   title: '状态',
  //   key: 'status',
  //   align: 'left',
  //   width: 80,
  //   render(row) {
  //     if (isNullObject(row.status)) {
  //       return ``;
  //     }
  //     return h(
  //       NTag,
  //       {
  //         style: {
  //           marginRight: '6px',
  //         },
  //         type: getOptionTag(options.value.sys_normal_disable, row.status),
  //         bordered: false,
  //       },
  //       {
  //         default: () => getOptionLabel(options.value.sys_normal_disable, row.status),
  //       }
  //     );
  //   },
  // },
  // {
  //   title: '排序',
  //   key: 'sort',
  //   align: 'left',
  //   width: 80,
  // },
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
    // for (const item of schemas.value) {
    //   switch (item.field) {
    //     case 'status':
    //       item.componentProps.options = options.value.sys_normal_disable;
    //       break;
    //   }
    // }
  });
}


