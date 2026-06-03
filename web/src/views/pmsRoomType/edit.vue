<template>
  <div>

    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }" :header-style="{
                    padding: '20px',
                  }" closable>
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">房型编辑
          </div>
        </template>
        <template #footer>
          <n-button style="width: 70px;height: 35px;margin-right: 10px" @click="closeForm">
            取消
          </n-button>
          <n-button :loading="formBtnLoading" style="width: 70px;height: 35px;" type="info"
                    @click="confirmForm">
            保存
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            label-placement="top"
            label-width="auto"
          >
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                房型展示图
              </div>
              <n-grid :cols="1">
                <n-gi>
                  <n-form-item label="封面图" path="cover" :show-feedback='false'>
                    <FileChooser1 v-model:value="formValue.cover" :maxNumber="1"
                                  fileType="default"/>
<!--                    <UploadImage :maxNumber="1" v-model:value="formValue.cover" />-->
                  </n-form-item>
                  <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 10px 0 24px">建议尺寸：150px*150px</div>
                </n-gi>
                <n-gi>
                  <n-form-item path="coverList" :show-feedback='false'>
                    <template #label>
                      <div style="display: flex;align-items: center;">
                        <div>图集</div>
                        <div style="margin-left: 10px;display: flex;align-items: center;">
                          <FileChooser
                            :maxNumber="10"
                            v-model:value="formValue.coverList"
                            :nolist="true"
                            @change="FileChooserchange"
                          />
                        </div>
                      </div>
                    </template>
                    <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px" v-if="formValue.coverList.length <= 0">建议尺寸：750px*420px</div>
                    <Draggable
                      animation="300"
                      :list="formValue.coverList"
                      group="people"
                      itemKey="fileUrl"
                      @update="dragchange"
                    >
                      <template #item="{ element }">
                        <div class="mr-3 mb-3" style="display: inline-block;width: 100px;">
                          <n-image-group>
                            <n-space>
                              <n-image object-fit="cover" :width="100" :src="element" style="height: 100px;border-radius: 5px;"/>
                            </n-space>
                          </n-image-group>
                          <div class="c333 flex-row f12">
                            <div class="text-r c999">
                              <a class="mr-2 c999" @click="delimg(element)"> 删除 </a>
                            </div>
                          </div>
                        </div>
                      </template>
                    </Draggable>
                  </n-form-item>
                  <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px" v-if="formValue.coverList.length > 0">建议尺寸：750px*420px</div>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                房型名称
              </div>
              <n-grid :cols="2" x-gap="10">
                <n-gi>
                  <n-form-item label="房型名称_简体中文" path="name_zh">
                    <n-input v-model:value="nameLanguage.zh" placeholder="简体中文房型名称"
                             style="width: 100%"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="房型名称_日本语" path="name_ja">
                    <n-input v-model:value="nameLanguage.ja" placeholder="日本语房型名称"
                             style="width: 100%"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="房型名称_English" path="name_en">
                    <n-input v-model:value="nameLanguage.en" placeholder="英语房型名称"
                             style="width: 100%"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="房型名称_한국어" path="name_ko">
                    <n-input v-model:value="nameLanguage.ko" placeholder="韩语房型名称"
                             style="width: 100%"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="房型名称_繁体中文" path="name_zh_CN">
                    <n-input v-model:value="nameLanguage.zh_CN" placeholder="繁体中文房型名称"
                             style="width: 100%"/>
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                费率计划
              </div>
              <n-grid :cols="1" :x-gap="24">
                <n-gi>
                  <n-form-item label="费率计划" path="bedrooms">


                    <n-radio-group v-model:value="formValue.ratePlanId" name="radiogroup">
                      <n-space>
                        <div style="display: flex;flex-direction: column;">
                          <n-radio v-for="Plan in formValue.roomRatePlanInfos" :key="Plan.id"
                                   v-model:value="Plan.ratePlanId">
                            {{ Plan.rateName }}
                          </n-radio>
                        </div>
                      </n-space>
                    </n-radio-group>


                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                房型参数
              </div>
              <n-grid :cols="2" x-gap="10">
                <n-gi>
                  <n-form-item label="房型间数（单位：间）" path="bedrooms">
                    <n-input-number v-model:value="formValue.roomNum" placeholder="请输入间数"
                                    style="width: 100%"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="最大容纳人数（单位：人）" path="occupancy">
                    <n-input-number v-model:value="formValue.occupancy" placeholder="请输入占用"
                                    style="width: 100%"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item path="checkinAt">
                    <template #label>
                      <span class="n-form-item-label__text">入住时间</span><span
                      class="n-form-item-label__asterisk">*</span><span
                      class="n-form-item-label__text">（GMT 9+）</span>
                    </template>
                    <n-time-picker v-model:formatted-value="formValue.checkinAt" format="HH:mm"
                                   style="width: 100%"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item path="checkoutAt">
                    <template #label>
                      <span class="n-form-item-label__text">退房时间</span><span
                      class="n-form-item-label__asterisk">*</span><span
                      class="n-form-item-label__text">（GMT 9+）</span>
                    </template>
                    <n-time-picker v-model:formatted-value="formValue.checkoutAt" format="HH:mm"
                                   style="width: 100%"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="房型面积（单位：m²）" path="size">
                    <n-input v-model:value="formValue.size" placeholder="请输入面积"
                             style="width: 100%"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="卧室数" path="bedrooms">
                    <n-input v-model:value="formValue.bedrooms" placeholder="请输入卧室数"
                             style="width: 100%"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="最低价格（单位：JPY）" path="basePrice">
                    <n-input-number v-model:value="formValue.basePrice" placeholder="请输入最低价格"
                                    style="width: 100%"/>
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <div style="width: 100%;display: flex;justify-content: space-between">
                    <div>
                      <n-form-item :show-feedback="false" label="价格人数（单位：人）"
                                   path="occupantsForBaseRate">
                        <n-input-number v-model:value="formValue.occupantsForBaseRate"
                                        :precision="0" placeholder="请输入人数"
                                        style="width: 120px"/>
                      </n-form-item>
                    </div>
                    <div>
                      <n-form-item :show-feedback="false" label="超员增加费用" path="appversion">
                        <a-radio-group v-model:value="isAddFee" name="isAddFee">
                          <a-radio value="1">增加</a-radio>
                          <a-radio style="margin-right: 0" value="2">不增加</a-radio>
                        </a-radio-group>
                      </n-form-item>
                    </div>
                  </div>
                </n-gi>
                <n-gi>
                  <n-form-item v-if="isAddFee == '1'" :show-feedback="false"
                               label="超员费用（单位：JPY）" path="additionalGuestAmounts">
                    <n-input-number v-model:value="formValue.additionalGuestAmounts"
                                    placeholder="请输入超员费用" style="width: 100%;"/>
                  </n-form-item>
                </n-gi>
                <n-gi span="2">
                  <div class="tips">
                    *价格人数：指的是房价包含的人数，如开启超员增加费用，超过此人数需设置额外加收费用，不开启则不按人数增加费用
                  </div>
                </n-gi>
              </n-grid>
            </div>
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                床型
              </div>
              <div v-for="(item, index) in bedList" :key="index" class="room-type-list">
                <n-grid :cols="2" x-gap="10">
                  <n-gi>
                    <div class="room-type-list-label">床型名称</div>
                    <n-input v-model:value="item.bedTypeName" placeholder="请输入床型名称"
                             style="width: 100%"/>
                  </n-gi>
                  <n-gi>
                    <div class="room-type-list-item">
                      <div class="room-type-list-item1">
                        <div class="room-type-list-label">床宽（单位：米）</div>
                        <n-input v-model:value="item.bedWidth" placeholder="请输入床宽"
                                 style="width: 100%"/>
                      </div>
                      <div class="room-type-list-item2">
                        <div class="room-type-list-label">数量（单位：个）</div>
                        <n-input v-model:value="item.bedNum" placeholder="请输入数量"
                                 style="width: 100%"/>
                      </div>
                      <n-button
                        :style="buttonCSS"
                        text
                        @click="handleDelBed(index)"
                      >
                        <n-icon size="18">
                          <TrashOutline/>
                        </n-icon>
                      </n-button>
                    </div>
                  </n-gi>
                </n-grid>
              </div>
              <n-button type="primary" @click="handleAddBed">
                <template #icon>
                  <n-icon>
                    <AddCircleOutline/>
                  </n-icon>
                </template>
                添加
              </n-button>
            </div>
          </n-form>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue';
import {Edit, View} from '@/api/pmsRoomType';
import {newState, rules, State} from './model';
import {useProjectSettingStore} from '@/store/modules/projectSetting';
import {NButton, NIcon, useMessage} from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';
import {jsontoobj} from "@/utils/smjcomm";
import {AddCircleOutline, TrashOutline} from "@vicons/ionicons5";
import Draggable from "vuedraggable";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const isAddFee = ref('1')
const bedList = ref([
  {
    bedTypeName: '',
    bedWidth: null,
    bedNum: null
  }
])
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const nameLanguage = ref({
  en: '',
  zh: '',
  ja: '',
  ko: '',
  zh_CN: '',
});

const buttonCSS = computed(() => {
  return {
    '--n-text-color': 'rgb(61, 61, 61)',
    '--n-text-color-hover': 'rgb(61, 61, 61)',
    '--n-text-color-pressed': 'rgb(61, 61, 61)',
    '--n-text-color-focus': 'rgb(61, 61, 61)',
    'margin-top': '40px'
  }
})

const delimg = (url) => {
  formValue.value.coverList = formValue.value.coverList
    .filter((f) => {
      if (f != url) {
        return f;
      }
    })
    .map((m) => {
      return m;
    });
};

const dragchange = () => {
  formValue.value.coverList = formValue.value.coverList.map((m) => {
    return m;
  });
};

const FileChooserchange = (res) => {
  console.log('res',res);
  formValue.value.coverList.concat(res.allvalue.map((m) => {
    return m.fileUrl;
  }));
};

function openModal(state: State) {
  showModal.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);

    return;
  }

  // 编辑
  loading.value = true;
  View({id: state.id})
    .then((res) => {
      res.bedrooms = parseInt(res.bedrooms)
      if (res.bedTypes) {
        bedList.value = res.bedTypes
      } else {
        bedList.value = [
          {
            bedTypeName: '',
            bedWidth: null,
            bedNum: null
          }
        ]
      }

      if (res.additionalGuestAmounts > 0) {
        isAddFee.value = "1"
      } else {
        isAddFee.value = "2"
      }

      console.log('isAddFee', isAddFee.value)

      res.bedrooms = res.bedrooms ? res.bedrooms : ''

      res.checkinAt = res.checkinAt ? res.checkinAt : '00:00';
      res.checkoutAt = res.checkoutAt ? res.checkoutAt : '00:00';

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
    })
    .finally(() => {
      loading.value = false;
    });
}

function confirmForm(e) {
  // console.log('bedList', bedList)
  formValue.value.bedType = bedList.value
  // console.log('formValue.value', formValue.value)
  // return
  formValue.value.nameLanguage = nameLanguage.value

  if (isAddFee.value == "2") {
    formValue.value.additionalGuestAmounts = 0;
  }

  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {

      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          formBtnLoading.value = false;
          closeForm();
          emit('reloadTable');
        });
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      message.error('请填写完整信息');
      formBtnLoading.value = false;
    }
  });
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
  bedList.value = [
    {
      bedTypeName: '',
      bedWidth: null,
      bedNum: null
    }
  ]
}

function handleAddBed() {
  let item = {
    bedTypeName: '',
    bedWidth: null,
    bedNum: null
  }
  bedList.value.push(item)
}

function handleDelBed(index) {
  bedList.value.splice(index, 1)
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

    &-content {

    }
  }
}

.tips {
  font-weight: 400;
  font-size: 14px;
  color: #F73314;
  line-height: 20px;
  margin-top: 8px;
  margin-bottom: 20px;
}

.room-type-list {
  margin-bottom: 20px;

  .room-type-list-label {
    padding: 0 0 6px 2px;
    font-size: 14px;
    line-height: 26px;
  }

  .room-type-list-item {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;

    .room-type-list-item1 {
      flex: 1;
      margin-right: 10px;
    }

    .room-type-list-item2 {
      flex: 1;
      margin-right: 10px;
    }
  }
}
</style>


