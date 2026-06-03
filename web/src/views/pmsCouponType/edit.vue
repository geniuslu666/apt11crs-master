<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
        padding: '20px',
      }" :body-content-style="{
        padding: '20px',
      }" :footer-style="{
        padding: '12px 20px',
      }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">
            {{ formValue.id > 0 ? '编辑优惠券#' + formValue.id : '添加优惠券' }}
          </div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;">
            保存
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-form ref="formRef" :model="formValue" :rules="rules" label-placement="top" label-width="auto">
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                优惠券名称
              </div>
              <n-grid cols="1 600:3" x-gap="80">
                <n-gi>
                  <n-form-item label="优惠券名称_简体中文" path="name_zh">
                    <n-input placeholder="简体中文优惠券名称" v-model:value="nameLanguage.zh" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="优惠券名称_日本语" path="name_ja">
                    <n-input placeholder="日本语优惠券名称" v-model:value="nameLanguage.ja" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="优惠券名称_English" path="name_en">
                    <n-input placeholder="英语优惠券名称" v-model:value="nameLanguage.en" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="优惠券名称_한국어" path="name_ko">
                    <n-input placeholder="韩语优惠券名称" v-model:value="nameLanguage.ko" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="优惠券名称_繁体中文" path="name_zh_CN">
                    <n-input placeholder="繁体中文优惠券名称" v-model:value="nameLanguage.zh_CN" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                基本信息
              </div>
              <n-grid cols="1 600:3" x-gap="80">
                <n-gi span="3">
                  <n-form-item label="优惠券类型" path="type">
                    <n-radio-group v-model:value="formValue.type">
                      <n-radio-button value="reward" label="满减" />
                      <n-radio-button value="discount" label="折扣" />
                    </n-radio-group>
                  </n-form-item>
                </n-gi>
                <n-gi v-if="formValue.type == 'reward'">
                  <n-form-item path="money" :show-require-mark="false">
                    <template #label>
                      <span class="n-form-item-label__text">优惠券面额</span><span
                        class="n-form-item-label__asterisk">*</span><span
                        class="n-form-item-label__text">（单位：JPY）</span>
                    </template>
                    <n-input-number :min="0" placeholder="优惠券面额" v-model:value="formValue.money" style="width: 300px" />
                  </n-form-item>
                </n-gi>
                <n-gi v-if="formValue.type == 'discount'">
                  <n-form-item path="discount" :show-require-mark="false">
                    <template #label>
                      <span class="n-form-item-label__text">优惠券折扣</span><span
                        class="n-form-item-label__asterisk">*</span><span class="n-form-item-label__text">（单位：折）</span>
                    </template>
                    <n-input-number :min="1" :max="9.9" placeholder="优惠券折扣" v-model:value="formValue.discount"
                      style="width: 300px" />
                  </n-form-item>
                </n-gi>
                <n-gi v-if="formValue.type == 'discount'">
                  <n-form-item path="discountLimit" :show-require-mark="false">
                    <template #label>
                      <span class="n-form-item-label__text">最多优惠</span><span
                        class="n-form-item-label__asterisk">*</span><span
                        class="n-form-item-label__text">（单位：JPY）</span>
                    </template>
                    <n-input-number :min="0" placeholder="最多优惠" v-model:value="formValue.discountLimit"
                      style="width: 300px" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item path="atLeast" :show-require-mark="false">
                    <template #label>
                      <span class="n-form-item-label__text">满多少可以使用</span><span
                        class="n-form-item-label__asterisk">*</span><span
                        class="n-form-item-label__text">（单位：JPY）</span>
                    </template>
                    <n-input-number :min="0" placeholder="请输入" v-model:value="formValue.atLeast" style="width: 300px" />
                  </n-form-item>
                </n-gi>
                <n-gi :span="3">
                  <n-form-item path="count" :show-require-mark="false">
                    <template #label>
                      <span class="n-form-item-label__text">发放数量</span><span
                        class="n-form-item-label__asterisk">*</span><span class="n-form-item-label__text">（单位：张）</span>
                    </template>
                    <n-input-number placeholder="发放数量" v-model:value="formValue.count" style="width: 300px" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="是否允许直接领取" path="isShow">
                    <n-switch v-model:value="formValue.isShow" :unchecked-value="0" :checked-value="1" />
                  </n-form-item>
                </n-gi>
                <n-gi v-if="formValue.isShow == 1">
                  <n-form-item path="maxFetch" :show-require-mark="false">
                    <template #label>
                      <span class="n-form-item-label__text">最大领取数量</span><span
                        class="n-form-item-label__asterisk">*</span><span class="n-form-item-label__text">（单位：张）</span>
                    </template>
                    <n-input-number :min="0" placeholder="最大领取数量" v-model:value="formValue.maxFetch"
                      style="width: 300px" />
                  </n-form-item>
                </n-gi>
                <n-gi span="3">
                  <n-form-item label="有效期类型" path="validityType">
                    <n-radio-group v-model:value="formValue.validityType">
                      <n-space>
                        <n-radio :value="1">
                          固定时间
                        </n-radio>
                        <n-radio :value="2">
                          发放/领取之日起
                        </n-radio>
                        <n-radio :value="3">
                          长期有效
                        </n-radio>
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="3" v-if="formValue.validityType == 1">
                  <n-form-item label="固定时间" path="endUseTime">
                    <DatePicker v-model:formValue="formValue.endUseTime" type="date" style="width: 300px"
                      :is-date-disabled="disablePreviousDate" />
                  </n-form-item>
                </n-gi>
                <n-gi span="3" v-if="formValue.validityType == 2">
                  <n-form-item path="fixedTerm" :show-require-mark="false">
                    <template #label>
                      <span class="n-form-item-label__text">发放/领取之日起</span><span
                        class="n-form-item-label__asterisk">*</span><span
                        class="n-form-item-label__text">（单位：天有效）</span>
                    </template>
                    <n-input-number :min="0" placeholder="发放/领取之日起" v-model:value="formValue.fixedTerm"
                      style="width: 300px" />
                  </n-form-item>
                </n-gi>
                <n-gi span="3">
                  <n-form-item label="使用场景" path="scene">
                    <n-radio-group v-model:value="formValue.scene">
                      <n-space>
                        <n-radio :value="1">
                          住宿
                        </n-radio>
                        <n-radio :value="2">
                          订餐
                        </n-radio>
                        <n-radio :value="3">
                          按摩
                        </n-radio>
                        <n-radio :value="4">
                          接送机/包车
                        </n-radio>
                        <n-radio :value="5">
                          储物柜
                        </n-radio>
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="3" v-if="formValue.scene == 1">
                  <n-form-item :show-feedback="false" label="物业" path="endIdsType">
                    <n-radio-group v-model:value="formValue.propertySelectType" name="propertySelectType">
                      <n-space>
                        <n-radio :value="1">
                          全部物业
                        </n-radio>
                        <n-radio :value="2">
                          部分物业
                        </n-radio>
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                  <n-form-item style="display: block;margin-top: 10px" path="propertyIdsArr"
                    v-if="formValue.propertySelectType == 2" :show-feedback="false">
                    <n-select placeholder="请选择物业" v-model:value="formValue.propertyIdsArr" :options="propertyList"
                      label-field="name" value-field="id" clearable filterable multiple style="width: 320px" />
                  </n-form-item>
                </n-gi>
                <n-gi span="3" v-if="formValue.scene == 2">
                  <n-form-item :show-feedback="false" label="餐厅" path="endIdsType">
                    <n-radio-group v-model:value="formValue.restaurantSelectType" name="restaurantSelectType">
                      <n-space>
                        <n-radio :value="1">
                          全部餐厅
                        </n-radio>
                        <n-radio :value="2">
                          部分餐厅
                        </n-radio>
                      </n-space>
                    </n-radio-group>
                  </n-form-item>
                  <n-form-item style="display: block;margin-top: 10px" path="restaurantIdsArr"
                    v-if="formValue.restaurantSelectType == 2" :show-feedback="false">
                    <n-select placeholder="请选择餐厅" v-model:value="formValue.restaurantIdsArr" :options="restaurantList"
                      label-field="name" value-field="id" clearable filterable multiple style="width: 320px" />
                  </n-form-item>
                </n-gi>
                <n-gi span="3" v-if="formValue.scene == 3">
                  <n-form-item label="服务类型" path="serviceIdsArr" :show-require-mark="true" :show-feedback="false">
                    <n-checkbox-group v-model:value="formValue.serviceIdsArr">
                      <n-space>
                        <n-checkbox value="ToStore" label="到店" />
                        <n-checkbox value="ToDoor" label="上门" />
                      </n-space>
                    </n-checkbox-group>
                  </n-form-item>
                </n-gi>
                <n-gi span="3" v-if="formValue.scene == 4">
                  <n-form-item label="服务类型" path="carServiceTypesArr" :show-require-mark="true" :show-feedback="false">
                    <n-checkbox-group v-model:value="formValue.carServiceTypesArr">
                      <n-space>
                        <n-checkbox value="PICKUP" label="接机" />
                        <n-checkbox value="DELIVERY" label="送机" />
                        <n-checkbox value="CHARTERED" label="包车" />
                      </n-space>
                    </n-checkbox-group>
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div" style="margin-top: 24px">
              <div class="level-detail-div-title">
                <div></div>
                使用说明
              </div>
              <n-grid cols="1 600:3" x-gap="80">
                <n-gi>
                  <n-form-item label="使用说明_简体中文" path="desc_zh">
                    <n-input type="textarea" placeholder="请输入简体中文使用说明" v-model:value="descLanguage.zh"
                      style="width: 100%" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="使用说明_日本语" path="desc_ja">
                    <n-input type="textarea" placeholder="请输入日本语使用说明" v-model:value="descLanguage.ja"
                      style="width: 100%" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="使用说明_English" path="desc_en">
                    <n-input type="textarea" placeholder="请输入英语使用说明" v-model:value="descLanguage.en"
                      style="width: 100%" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="使用说明_한국어" path="desc_ko">
                    <n-input type="textarea" placeholder="请输入韩语使用说明" v-model:value="descLanguage.ko"
                      style="width: 100%" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="使用说明_繁体中文" path="desc_zh_CN">
                    <n-input type="textarea" placeholder="请输入繁体中文使用说明" v-model:value="descLanguage.zh_CN"
                      style="width: 100%" />
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { Edit, View } from '@/api/pmsCouponType';
import { State, newState } from './model';
import DatePicker from '@/components/DatePicker/datePicker.vue';
import { FormItemRule, useMessage } from 'naive-ui';
import { All } from "@/api/pmsProperty";
import { All as AllRestaurant } from "@/api/foodRestaurant";
import { jsontoobj } from "@/utils/smjcomm";
import { subDays } from "date-fns/esm";
import { adaModalWidth } from "@/utils/hotgo";

const emit = defineEmits(['reloadTable']);
const loading = ref(false);
const showModal = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(1100);
});

const propertyList = ref([]);
const restaurantList = ref([]);
const formValue = ref<State>(newState(null));

const nameLanguage = ref({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

const descLanguage = ref({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

// 表单验证规则
const rules = {
  money: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: number) {
      if (formValue.value.type == 'reward' && (value == null || value < 0)) {
        return new Error('请输入优惠券面额')
      }
      return true
    },
  },
  discount: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: number) {
      if (formValue.value.type == 'discount' && !value) {
        return new Error('请输入优惠券折扣')
      }
      return true
    },
  },
  discountLimit: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: number) {
      if (value == null || value < 0) {
        return new Error('请输入最多优惠')
      }
      return true
    },
  },
  atLeast: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: number) {
      if (value == null || value < 0) {
        return new Error('请输入满多少JPY可以使用')
      }
      return true
    },
  },
  count: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: number) {
      if (value == null) {
        return new Error('请输入发放数量')
      }
      return true
    },
  },
  maxFetch: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: number) {
      if (value == null || value < 0) {
        return new Error('请输入最大领取数量')
      }
      return true
    },
  },
  endUseTime: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: number) {
      if (formValue.value.validityType == 1 && !value) {
        return new Error('请选择固定时间')
      }
      return true
    },
  },
  fixedTerm: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: number) {
      if (formValue.value.validityType == 2 && (value == null || value < 0)) {
        return new Error('请输入领取之日起多少天有效')
      }
      return true
    },
  },
  /*propertyIdsArr: {
    required: true,
    trigger: ['blur', 'input'],
    validator(rule: FormItemRule, value: string) {
      console.log(value)
      if (formValue.value.scene == 1 && (!value || value.length <= 0)) {
        return new Error('请选择物业')
      }
      return true
    },
  },*/
};
const message = useMessage();
const formRef = ref<any>({});
const formBtnLoading = ref(false);

function disablePreviousDate(ts: number) {
  return ts < subDays(Date.now(), 1).valueOf();
}

function confirmForm(e) {
  // 优惠券名称-多语言
  formValue.value.nameLanguage = nameLanguage.value
  // 优惠券使用说明-多语言
  formValue.value.descLanguage = descLanguage.value

  formValue.value.couponName = ""
  formValue.value.desc = ""

  console.log('formValue', formValue.value)

  if (formValue.value.scene == 1 && formValue.value.propertySelectType == 2 && formValue.value.propertyIdsArr.length <= 0) {
    formBtnLoading.value = false;
    message.error('请选择物业');
    return false;
  }
  if (formValue.value.scene == 2 && formValue.value.restaurantSelectType == 2 && formValue.value.restaurantIdsArr.length <= 0) {
    formBtnLoading.value = false;
    message.error('请选择餐厅');
    return false;
  }
  // return

  e.preventDefault();
  formValue.value.propertyIds = formValue.value.propertyIdsArr ? formValue.value.propertyIdsArr.join(',') : ''
  formValue.value.restaurantIds = formValue.value.restaurantIdsArr ? formValue.value.restaurantIdsArr.join(',') : ''
  formValue.value.carServiceTypes = formValue.value.carServiceTypesArr ? formValue.value.carServiceTypesArr.join(',') : ''

  if (formValue.value.scene == 1 && formValue.value.propertySelectType == 1) {
    formValue.value.propertyIds = '';
  }
  if (formValue.value.scene == 2 && formValue.value.restaurantSelectType == 1) {
    formValue.value.restaurantIds = '';
  }


  // 不允许直接领取(那么不限制发放数量，和最大领取数量)
  if (formValue.value.isShow == 0) {
    formValue.value.count = 0;
    formValue.value.maxFetch = 0
  }

  if (formValue.value.scene == 3 && formValue.value.serviceIdsArr.length <= 0) {
    message.error('请选择按摩服务类型');
    formBtnLoading.value = false;
    return false;
  }

  if (formValue.value.scene == 4 && formValue.value.carServiceTypesArr.length <= 0) {
    message.error('请选择服务类型');
    formBtnLoading.value = false;
    return false;
  }

  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        closeForm();
        emit('reloadTable');
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      message.error('请填写完整信息');
      formBtnLoading.value = false;
    }
  });
}

async function getAllProperty() {
  const res = await All({});
  propertyList.value = res.list;
}

async function getAllRestaurant() {
  const res = await AllRestaurant({});
  restaurantList.value = res.list;
}

async function getInfo(id) {
  const res = await View({ id: id });
  formValue.value = res;

  formValue.value.serviceIdsArr = res.serviceIds ? res.serviceIds.split(",").map(item => String(item)) : [];
  formValue.value.carServiceTypesArr = res.carServiceTypes ? res.carServiceTypes.split(",").map(item => String(item)) : [];

  // 住宿
  if (res.propertyIds) {
    formValue.value.propertyIdsArr = res.propertyIds ? res.propertyIds.split(",").map(item => Number(item)) : [];
    formValue.value.propertySelectType = 2
  } else {
    formValue.value.propertyIdsArr = [];
    formValue.value.propertySelectType = 1
  }

  // 餐厅
  if (res.restaurantIds) {
    formValue.value.restaurantIdsArr = res.restaurantIds ? res.restaurantIds.split(",").map(item => Number(item)) : [];
    formValue.value.restaurantSelectType = 2
  } else {
    formValue.value.restaurantIdsArr = [];
    formValue.value.restaurantSelectType = 1
  }

  if (res.nameLanguage) {
    res.nameLanguage = jsontoobj(res.nameLanguage);
    nameLanguage.value.zh = res.nameLanguage.zh && res.nameLanguage.zh.content ? res.nameLanguage.zh.content : ''
    nameLanguage.value.en = res.nameLanguage.en && res.nameLanguage.en.content ? res.nameLanguage.en.content : ''
    nameLanguage.value.ko = res.nameLanguage.ko && res.nameLanguage.ko.content ? res.nameLanguage.ko.content : ''
    nameLanguage.value.ja = res.nameLanguage.ja && res.nameLanguage.ja.content ? res.nameLanguage.ja.content : ''
    if (res.nameLanguage.zh_CN) {
      nameLanguage.value.zh_CN = res.nameLanguage.zh_CN.content ? res.nameLanguage.zh_CN.content : ''
    } else if (res.nameLanguage.zh_cn) {
      nameLanguage.value.zh_CN = res.nameLanguage.zh_cn.content ? res.nameLanguage.zh_cn.content : ''
    } else {
      nameLanguage.value.zh_CN = ''
    }
  }

  if (res.descLanguage) {
    res.descLanguage = jsontoobj(res.descLanguage);
    descLanguage.value.zh = res.descLanguage.zh && res.descLanguage.zh.content ? res.descLanguage.zh.content : ''
    descLanguage.value.en = res.descLanguage.en && res.descLanguage.en.content ? res.descLanguage.en.content : ''
    descLanguage.value.ko = res.descLanguage.ko && res.descLanguage.ko.content ? res.descLanguage.ko.content : ''
    descLanguage.value.ja = res.descLanguage.ja && res.descLanguage.ja.content ? res.descLanguage.ja.content : ''
    if (res.descLanguage.zh_CN) {
      descLanguage.value.zh_CN = res.descLanguage.zh_CN.content ? res.descLanguage.zh_CN.content : ''
    } else if (res.descLanguage.zh_cn) {
      descLanguage.value.zh_CN = res.descLanguage.zh_cn.content ? res.descLanguage.zh_cn.content : ''
    } else {
      descLanguage.value.zh_CN = ''
    }
  }
}


async function openModal(state: State) {
  showModal.value = true;
  loading.value = true;

  await getAllProperty();
  await getAllRestaurant();

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    nameLanguage.value = {
      zh: '',
      en: '',
      ko: '',
      ja: '',
      zh_CN: '',
    };
    descLanguage.value = {
      zh: '',
      en: '',
      ko: '',
      ja: '',
      zh_CN: '',
    };

    loading.value = false;

    return;
  }

  await getInfo(state.id)
  loading.value = false;

}

function closeForm() {
  showModal.value = false;
  loading.value = false;
}

defineExpose({
  openModal,
});

</script>

<style lang="less">
.level-detail-div {
  &-title {
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 15px;

    div {
      margin-right: 5px;
      width: 6px;
      height: 15px;
      background: #053DC8;
    }
  }

  &-item {
    &-title {
      font-weight: 400;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      margin-bottom: 8px;
    }

    &-content {}
  }
}
</style>
