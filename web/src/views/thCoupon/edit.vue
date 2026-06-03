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
            {{ formValue.id > 0 ? '编辑礼品券#' + formValue.id : '添加礼品券' }}
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
                礼品券名称
              </div>
              <n-grid cols="1 600:3" x-gap="80">
                <n-gi>
                  <n-form-item label="礼品券名称_简体中文" path="name_zh" :show-require-mark="false">
                    <n-input placeholder="简体中文礼品券名称" v-model:value="nameLanguage.zh" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="礼品券名称_English" path="name_en" :show-require-mark="false">
                    <n-input placeholder="英语礼品券名称" v-model:value="nameLanguage.en" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="礼品券名称_日本语" path="name_ja" :show-require-mark="false">
                    <n-input placeholder="日本语礼品券名称" v-model:value="nameLanguage.ja" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>

                <n-gi>
                  <n-form-item label="副标题_简体中文" path="sub_name_zh" :show-require-mark="false">
                    <n-input placeholder="简体中文副标题" v-model:value="subNameLanguage.zh" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="副标题_English" path="sub_name_en" :show-require-mark="false">
                    <n-input placeholder="英语副标题" v-model:value="subNameLanguage.en" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="副标题_日本语" path="sub_name_ja" :show-require-mark="false">
                    <n-input placeholder="日本语副标题" v-model:value="subNameLanguage.ja" :style="{ width: '300px' }" />
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
                  <n-form-item label="礼品券分类" path="categoryId" :show-require-mark="false">
                    <n-select v-model:value="formValue.categoryId" :options="cateList" clearable filterable
                      label-field="name" value-field="id" style="width: 300px" />
                  </n-form-item>
                </n-gi>
                <n-gi span="3">
                  <n-form-item label="券识别名称" path="couponNoPrefix">
                    <n-input placeholder="请输入券识别名称" v-model:value="formValue.identityName" style="width: 300px"/>
                  </n-form-item>
                </n-gi>
                <n-gi span="3">
                  <n-form-item label="编号前缀" path="couponNoPrefix">
                    <n-input placeholder="请输入编号前缀" v-model:value="formValue.couponNoPrefix" style="width: 300px" />
                  </n-form-item>
                </n-gi>
<!--                <n-gi>
                  <n-form-item label="券LOGO" path="logo" :show-feedback='false'>
                    <UploadImage :maxNumber="1" v-model:value="formValue.logo" />
                  </n-form-item>
                  <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px">建议尺寸：238px*270px</div>
                </n-gi>-->
                <n-gi span="3">
                  <n-form-item label="是否启用发行" path="status">
                    <n-switch v-model:value="formValue.status" :unchecked-value="2" :checked-value="1" />
                  </n-form-item>
                </n-gi>
                <n-gi span="3">
                  <n-form-item label="使用模式" path="useMode">
                    <n-radio-group v-model:value="formValue.useMode">
                      <n-radio-button value="ARRIVE_VERIFY" label="到店核销" />
                    </n-radio-group>
                  </n-form-item>
                </n-gi>

                <n-gi span="3">
                  <n-form-item label="适用商户" path="mchIds" :show-require-mark="true">
                    <n-button type="primary" @click="chooseMch()">添加商户</n-button>
                  </n-form-item>
                </n-gi>
                <n-gi span="3">
                  <n-table style="margin-bottom: 25px">
                    <thead>
                      <tr>
                        <th>商户名称</th>
                        <th>商户分类</th>
                        <th>商户状态</th>
                        <th>商户核销商品名</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="(item, index) of selectMchArr" :key="index">
                        <td>{{ item.name }}</td>
                        <td>{{ item.thMchCategoryName }}</td>
                        <td>
                          <n-tag :type="getOptionTag(mchOptions.sys_normal_disable, item.status)" size="small"
                            class="min-left-space">
                            {{ getOptionLabel(mchOptions.sys_normal_disable, item.status) }}
                          </n-tag>
                        </td>
                        <td>
                          <n-input v-model:value="item.aliasName" placeholder="商户核销商品名" style="width: 200px" />
                        </td>
                        <td>
                          <n-button type="error" @click="handleDeleteMch(item.id)">删除</n-button>
                        </td>
                      </tr>
                    </tbody>
                  </n-table>
                </n-gi>

                <n-gi span="3">
                  <n-form-item path="fixedTerm" label="有效期">
                    <n-input-group>
                      <n-input-group-label>激活之日起</n-input-group-label>
                      <n-input-number placeholder="请输入" :min="0" :precision="0" v-model:value="formValue.fixedTerm"
                        style="width: 100px" />
                      <n-input-group-label>天有效</n-input-group-label>
                    </n-input-group>
                  </n-form-item>
                </n-gi>

                <n-gi span="3">
                  <n-form-item label="是否需要预约" path="needReservation">
                    <n-switch v-model:value="formValue.needReservation" :unchecked-value="0" :checked-value="1" />
                  </n-form-item>
                </n-gi>

                <n-gi span="3" v-if="formValue.needReservation == 1">
                  <n-form-item style="display: block;" path="reservationRestaurantIdsArr" label="需预约餐厅">
                    <n-select placeholder="请选择餐厅" v-model:value="formValue.reservationRestaurantIdsArr"
                      :options="restaurantList" label-field="name" value-field="id" clearable filterable multiple
                      style="width: 320px" />
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                使用说明
              </div>
              <n-grid :cols="1">
                <n-gi>
                  <n-form-item label="使用说明_简体中文" path="desc_zh" :show-require-mark="false">
                    <n-input type="textarea" placeholder="请输入简体中文使用说明" v-model:value="descLanguage.zh"
                      style="width: 100%" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="使用说明_English" path="desc_en" :show-require-mark="false">
                    <n-input type="textarea" placeholder="请输入英语使用说明" v-model:value="descLanguage.en"
                      style="width: 100%" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="使用说明_日本语" path="desc_ja" :show-require-mark="false">
                    <n-input type="textarea" placeholder="请输入日本语使用说明" v-model:value="descLanguage.ja"
                      style="width: 100%" />
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>

    <ChooseMch ref="chooseMchRef" @reloadMch="handleChooseMch" />
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { Edit, View, MchList } from '@/api/thCoupon';
import { State, newState, cateList } from './model';
import { useMessage } from 'naive-ui';
import { jsontoobj } from "@/utils/smjcomm";
import { adaModalWidth, getOptionLabel, getOptionTag } from "@/utils/hotgo";
import { Text } from "@/api/translate";
import { options as mchOptions, loadOptions as loadMchOptions } from "@/views/thMch/model";
import ChooseMch from "@/views/thCoupon/chooseMch.vue";
import { List } from "@/api/thMch";
import UploadImage from "@/components/Upload/uploadImage.vue";
import { All as AllRestaurant } from "@/api/foodRestaurant";

interface Merchant {
  id: number;
  name: string;
  thMchCategoryName: string;
  status: number;
  aliasName: string;
}

const emit = defineEmits(['reloadTable']);
const loading = ref(false);
const showModal = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(1100);
});
const selectMchInfoArr = ref<Merchant[]>([])
const selectMchArr = ref<Merchant[]>([])
const selectMchIdsArr = ref<number[]>([])
const chooseMchRef = ref();
const restaurantList = ref([]);

const formValue = ref<State>(newState(null));

const nameLanguage = ref({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

const subNameLanguage = ref({
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
  couponNoPrefix: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入编号前缀',
  },
  fixedTerm: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入有效期',
  },
};
const message = useMessage();
const formRef = ref<any>({});
const formBtnLoading = ref(false);

function translate(type, lang, text) {
  if (type == 'name') {
    if (lang == 'zh') {
      // 简体中文
      nameLanguage.value.zh_CN = ''
    } else if (lang == 'en') {
      // 韩语
      nameLanguage.value.ko = ''
    }
  } else if (type == 'sub_name') {
    if (lang == 'zh') {
      // 简体中文
      subNameLanguage.value.zh_CN = ''
    } else if (lang == 'en') {
      // 韩语
      subNameLanguage.value.ko = ''
    }
  } else if (type == 'desc') {
    if (lang == 'zh') {
      // 简体中文
      descLanguage.value.zh_CN = ''
    } else if (lang == 'en') {
      // 韩语
      descLanguage.value.ko = ''
    }
  }
  if (text) {
    if (lang == 'en') {
      text = text.toLowerCase()
    }
    Text({
      text: text
    }).then((_res) => {
      if (type == 'name') {
        if (lang == 'zh') {
          // 简体中文
          nameLanguage.value.zh_CN = _res.zh_CN
        } else if (lang == 'en') {
          // 韩语
          nameLanguage.value.ko = _res.ko
        }
      } else if (type == 'sub_name') {
        if (lang == 'zh') {
          // 简体中文
          subNameLanguage.value.zh_CN = _res.zh_CN
        } else if (lang == 'en') {
          // 韩语
          subNameLanguage.value.ko = _res.ko
        }
      } else if (type == 'desc') {
        if (lang == 'zh') {
          // 简体中文
          descLanguage.value.zh_CN = _res.zh_CN
        } else if (lang == 'en') {
          // 韩语
          descLanguage.value.ko = _res.ko
        }
      }
    }).catch((err) => {

    });
  }
}

function chooseMch() {
  chooseMchRef.value.openModal(selectMchIdsArr.value);
}

async function handleChooseMch(selectedIds) {
  selectMchIdsArr.value = selectedIds
  loading.value = true;
  await getMchList()

  const oldMap = new Map(selectMchArr.value.map(m => [m.id, m]))
  console.log('oldMap', oldMap)
  selectMchArr.value = selectMchInfoArr.value.map(m => ({
    ...m,
    aliasName: oldMap.get(m.id)?.aliasName || ''
  }))
  console.log('selectMchArr', selectMchArr.value)
  loading.value = false;
}

function handleDeleteMch(id) {
  selectMchArr.value = selectMchArr.value.filter(m => m.id !== id)
  selectMchIdsArr.value = selectMchIdsArr.value.filter(m => m !== id)
}

async function getAllRestaurant() {
  const res = await AllRestaurant({});
  restaurantList.value = res.list;
}

async function getMchList() {
  const res = await List({
    pagination: false,
    mchIds: selectMchIdsArr.value
  });
  selectMchInfoArr.value = res.list
  console.log('selectMchInfoArr', selectMchInfoArr.value)
}

function confirmForm(e) {
  nameLanguage.value.zh_CN = nameLanguage.value.ja
  nameLanguage.value.ko = nameLanguage.value.ja
  subNameLanguage.value.zh_CN = subNameLanguage.value.ja
  subNameLanguage.value.ko = subNameLanguage.value.ja
  descLanguage.value.zh_CN = descLanguage.value.ja
  descLanguage.value.ko = descLanguage.value.ja
  // 礼品券名称-多语言
  formValue.value.nameLanguage = nameLanguage.value
  // 礼品券副标题-多语言
  formValue.value.subNameLanguage = subNameLanguage.value
  // 礼品券使用说明-多语言
  formValue.value.descLanguage = descLanguage.value

  formValue.value.couponName = ""
  formValue.value.couponSubName = ""
  formValue.value.desc = ""

  console.log('formValue', formValue.value)

  if (!nameLanguage.value.zh || !nameLanguage.value.en || !nameLanguage.value.ja) {
    formBtnLoading.value = false;
    message.error('礼品券名称请填写完整');
    return false;
  }
  if (!subNameLanguage.value.zh || !subNameLanguage.value.en || !subNameLanguage.value.ja) {
    formBtnLoading.value = false;
    message.error('礼品券副标题请填写完整');
    return false;
  }
  if (!descLanguage.value.zh || !descLanguage.value.en || !descLanguage.value.ja) {
    formBtnLoading.value = false;
    message.error('礼品券使用说明请填写完整');
    return false;
  }

  if (!formValue.value.categoryId) {
    formBtnLoading.value = false;
    message.error('请选择礼品券分类');
    return false;
  }

  if (selectMchArr.value.length <= 0) {
    formBtnLoading.value = false;
    message.error('请选择适用商户');
    return false;
  }

  formValue.value.mchList = selectMchArr.value.map(m => ({
    mchId: m.id,
    name: m.aliasName
  }))

  if (formValue.value.needReservation == 1 && formValue.value.reservationRestaurantIdsArr.length <= 0) {
    formBtnLoading.value = false;
    message.error('请选择需要预约的餐厅');
    return false;
  }

  formValue.value.reservationRestaurantIds = formValue.value.reservationRestaurantIdsArr ? formValue.value.reservationRestaurantIdsArr.join(',') : ''

  // return

  e.preventDefault();

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


async function getInfo(id) {
  const res = await View({ id: id });
  formValue.value = res;

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

  if (res.subNameLanguage) {
    res.subNameLanguage = jsontoobj(res.subNameLanguage);
    subNameLanguage.value.zh = res.subNameLanguage.zh && res.subNameLanguage.zh.content ? res.subNameLanguage.zh.content : ''
    subNameLanguage.value.en = res.subNameLanguage.en && res.subNameLanguage.en.content ? res.subNameLanguage.en.content : ''
    subNameLanguage.value.ko = res.subNameLanguage.ko && res.subNameLanguage.ko.content ? res.subNameLanguage.ko.content : ''
    subNameLanguage.value.ja = res.subNameLanguage.ja && res.subNameLanguage.ja.content ? res.subNameLanguage.ja.content : ''
    if (res.subNameLanguage.zh_CN) {
      subNameLanguage.value.zh_CN = res.subNameLanguage.zh_CN.content ? res.subNameLanguage.zh_CN.content : ''
    } else if (res.subNameLanguage.zh_cn) {
      subNameLanguage.value.zh_CN = res.subNameLanguage.zh_cn.content ? res.subNameLanguage.zh_cn.content : ''
    } else {
      subNameLanguage.value.zh_CN = ''
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

  selectMchArr.value = []
  selectMchIdsArr.value = []
  if (res.mchList) {
    res.mchList.forEach((item) => {
      selectMchArr.value.push({
        id: item.mchId,
        name: item.mchInfo?.name,
        thMchCategoryName: item.mchInfo?.categoryInfo?.name,
        status: item.mchInfo?.status,
        aliasName: item.name,
      })
      selectMchIdsArr.value.push(item.mchId)
    })
  }

  if (res.reservationRestaurantIds) {
    formValue.value.reservationRestaurantIdsArr = res.reservationRestaurantIds ? res.reservationRestaurantIds.split(",").map(item => Number(item)) : [];
  } else {
    formValue.value.reservationRestaurantIdsArr = [];
  }
}

async function openModal(state: State) {
  showModal.value = true;
  loading.value = true;

  await loadMchOptions()
  await getAllRestaurant();

  // 新增
  if (!state || state.id < 1) {
    selectMchArr.value = []
    formValue.value = newState(state);
    nameLanguage.value = {
      zh: '',
      en: '',
      ko: '',
      ja: '',
      zh_CN: '',
    };
    subNameLanguage.value = {
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
