import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { validate } from '@/utils/validateUtil';
import {NImage} from "naive-ui";
import {errorImg, Option} from "@/utils/hotgo";
import { Dicts } from '@/api/dict/dict';

export class State {
  public id = 0; // id
  public level = 0; // 会员等级
  public language = null; // 语言
  public levelName = ''; // 会员等级名称
  public exp = 0; // 会员经验
  public desc = ''; // 等级说明
  public wordColor = ''; // 等级字体颜色
  public levelBadge = ''; // 等级徽章（单图）
  public levelCard = ''; // 等级卡片（单图）
  public levelBigPic = ''; // 等级大图（单图）
  public hotelGetRate = 1; // 酒店场景获取积分倍率
  public foodGetRate = 1; // 餐饮场景获取积分倍率
  public spaGetRate = 1; // 按摩场景获取积分倍率
  public carGetRate = 1; // 接送机/包车场景获取积分倍率
  public cabinetGetRate = 1; // 储物柜场景获取积分倍率
  public memberCount = 0; // 等级持有人数
  public createdAt = ''; // created_at
  public updatedAt = ''; // updated_at
  public deletedAt = ''; // deleted_at
  public status = 1; // 状态

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
  level: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入会员等级,且为正整数',
    validator: validate.num,
  },
  levelName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入会员等级名称',
  },
  exp: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入会员经验',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'levelName',
    component: 'NInput',
    label: '等级名称',
    componentProps: {
      placeholder: '请输入会员等级名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '等级id',
    key: 'id',
    align: 'left',
    width: 100,
  },
  {
    title: '等级名称',
    key: 'levelName',
    align: 'left',
    width: -1,
  },
  {
    title: '等级徽章',
    key: 'levelBadge',
    align: 'left',
    width: -1,
    render(row) {
      if(row.levelBadge){
        return h(NImage, {
          width: 32,
          height: 32,
          src: row.levelBadge,
          fallbackSrc: errorImg,
          onError: errorImg,
          style: {
            width: '32px',
            height: '32px',
            'max-width': '100%',
            'max-height': '100%',
          },
        });
      }else{
        return '--'
      }
    },
  },
  {
    title: '等级卡片',
    key: 'levelCard',
    align: 'left',
    width: -1,
    render(row) {
      if(row.levelCard){
        return h(NImage, {
          width: 32,
          height: 32,
          src: row.levelCard,
          fallbackSrc: errorImg,
          onError: errorImg,
          style: {
            width: '32px',
            height: '32px',
            'max-width': '100%',
            'max-height': '100%',
          },
        });
      }else{
        return '--'
      }
    },
  },
  {
    title: '等级大图',
    key: 'levelBigPic',
    align: 'left',
    width: -1,
    render(row) {
      if(row.levelBigPic){
        return h(NImage, {
          width: 32,
          height: 32,
          src: row.levelBigPic,
          fallbackSrc: errorImg,
          onError: errorImg,
          style: {
            width: '32px',
            height: '32px',
            'max-width': '100%',
            'max-height': '100%',
          },
        });
      }else{
        return '--'
      }
    },
  },
  {
    title: '达到经验值',
    key: 'exp',
    align: 'left',
    width: -1,
  },
  {
    title: '持有人数',
    key: 'memberCount',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 165,
  },
];

// 字典数据选项
export const options = ref({
  sys_normal_disable: [] as Option[],
  language: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable', 'language'],
  }).then((res) => {
    options.value = res;
    // console.log('res', res)
    for (const item of schemas.value) {
      switch (item.field) {
        case 'status':
          item.componentProps.options = options.value.sys_normal_disable;
          break;
        case 'language':
          item.componentProps.options = options.value.language;
          break;
      }
    }
  });
}
