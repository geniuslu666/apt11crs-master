<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content :body-content-style="{
                    padding: '20px 20px 0',
                  }" :header-style="{
                    padding: '20px',
                  }" closable>
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">房型详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div" style="margin-bottom: 15px">
            <div class="level-detail-div-title">
              <div></div>
              房型展示图
            </div>
            <n-grid :cols="1" y-gap="25">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">封面图</div>
                  <div class="level-detail-div-item-content">
                    <n-image v-if="formValue.cover" :src="formValue.cover" width="95px"/>
                    <div v-else class="no-image-div">
                      <img src="~@/assets/images/empty_image_icon.png" width="30"/>
                      <span>暂无图片</span>
                    </div>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">图集</div>
                  <div class="level-detail-div-item-content" style="flex-wrap: wrap">
                    <n-image-group v-if="formValue?.coverList.length > 0">
                       <template v-for="(item, key) in formValue?.coverList" :key="key">
                          <n-image :src="item" style="margin-right: 10px;margin-bottom: 10px;" width="95px"/>
                        </template>
                    </n-image-group>
                    <div v-else class="no-image-div">
                      <img src="~@/assets/images/empty_image_icon.png" width="30"/>
                      <span>暂无图片</span>
                    </div>
                  </div>
                </div>
              </n-gi>
            </n-grid>
          </div>
          <div class="level-detail-div">
            <div class="level-detail-div-title">
              <div></div>
              房型名称
            </div>
            <n-grid :cols="2" y-gap="25">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">简体中文</div>
                  <div class="level-detail-div-item-content">{{ nameLanguage.zh ? nameLanguage.zh : '-' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">日本语</div>
                  <div class="level-detail-div-item-content">{{ nameLanguage.ja ? nameLanguage.ja : '-' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">English</div>
                  <div class="level-detail-div-item-content">{{ nameLanguage.en ? nameLanguage.en : '-' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">한국어</div>
                  <div class="level-detail-div-item-content">{{ nameLanguage.ko ? nameLanguage.ko : '-' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">繁体中文</div>
                  <div class="level-detail-div-item-content">{{ nameLanguage.zh_CN ? nameLanguage.zh_CN : '-' }}</div>
                </div>
              </n-gi>
            </n-grid>
          </div>
          <div class="level-detail-div">
            <div class="level-detail-div-title">
              <div></div>
              房型参数
            </div>
            <n-grid :cols="2" y-gap="25">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">房型间数（单位：间）</div>
                  <div class="level-detail-div-item-content">{{ formValue.roomNum }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">最大容纳人数（单位：人）</div>
                  <div class="level-detail-div-item-content">{{ formValue.occupancy }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">入住时间（GMT 9+）</div>
                  <div class="level-detail-div-item-content">{{ formValue.checkinAt }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">退房时间（GMT 9+）</div>
                  <div class="level-detail-div-item-content">{{ formValue.checkoutAt }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">房型面积（单位：m²）</div>
                  <div class="level-detail-div-item-content">{{ formValue.size }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">卧室数量</div>
                  <div class="level-detail-div-item-content">{{ formValue.bedrooms }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">最低价格（单位：JPY）</div>
                  <div class="level-detail-div-item-content">{{ formValue.basePrice }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">无需增加额外金额客人数（单位：人）</div>
                  <div class="level-detail-div-item-content">{{ formValue.occupantsForBaseRate }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">额外客人金额（单位：JPY）</div>
                  <div class="level-detail-div-item-content">{{ formValue.additionalGuestAmounts }}</div>
                </div>
              </n-gi>
            </n-grid>
          </div>

          <div v-if="formValue?.bedType" class="level-detail-div">
            <div class="level-detail-div-title">
              <div></div>
              床型
            </div>
            <n-grid :cols="3" y-gap="25">
              <template v-for="(item, key) in formValue?.bedTypes" :key="key">
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">床型{{ key+1 }}</div>
                    <div class="level-detail-div-item-content">{{ item.bedTypeName }}</div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">床宽（单位：米）</div>
                    <div class="level-detail-div-item-content">{{ item.bedWidth }}</div>
                  </div>
                </n-gi>
                <n-gi>
                  <div class="level-detail-div-item">
                    <div class="level-detail-div-item-title">数量（单位：个）</div>
                    <div class="level-detail-div-item-content">{{ item.bedNum }}</div>
                  </div>
                </n-gi>
              </template>
            </n-grid>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue';
import {View} from '@/api/pmsRoomType';
import {newState, State} from './model';
import {adaModalWidth} from '@/utils/hotgo';
import {jsontoobj} from "@/utils/smjcomm";

const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const nameLanguage = ref({
    en: '',
    zh: '',
    ja: '',
    ko: '',
    zh_CN: '',
  });
  const dialogWidth = computed(() => {
    return adaModalWidth(580);
  });

  function openModal(state: State) {
    showModal.value = true;
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;

        if(res.nameLanguage){
          res.nameLanguage = jsontoobj(res.nameLanguage);
          nameLanguage.value.zh = res.nameLanguage.zh && res.nameLanguage.zh.content ? res.nameLanguage.zh.content : ''
          nameLanguage.value.en = res.nameLanguage.en && res.nameLanguage.en.content ? res.nameLanguage.en.content : ''
          nameLanguage.value.ko = res.nameLanguage.ko && res.nameLanguage.ko.content ? res.nameLanguage.ko.content : ''
          nameLanguage.value.ja = res.nameLanguage.ja && res.nameLanguage.ja.content ? res.nameLanguage.ja.content : ''
          if(res.nameLanguage.zh_CN){
            nameLanguage.value.zh_CN = res.nameLanguage.zh_CN.content ? res.nameLanguage.zh_CN.content : ''
          }else if(res.nameLanguage.zh_cn){
            nameLanguage.value.zh_CN = res.nameLanguage.zh_cn.content ? res.nameLanguage.zh_cn.content : ''
          }else{
            nameLanguage.value.zh_CN = ''
          }
        }
      })
      .finally(() => {
        loading.value = false;
      });
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less" scoped>
.level-detail-div{
  margin-bottom: 25px;
  &-title{
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 25px;
    div{
      margin-right: 5px;
      width: 6px;
      height: 15px;
      background: #053DC8;
    }
  }
  &-item{
    &-title{
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
      margin-bottom: 8px;
    }
    &-content{
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      display: flex;
      align-items: center;
      .color{
        width: 20px;
        height: 20px;
        margin-right: 8px;
        border-radius: 2px;
      }
    }
  }
}
.no-image-div{
  width: 95px;
  height: 95px;
  background: #FAFAFC;
  border: 1px solid #E0E0E6;
  border-radius: 2px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  span{
    font-weight: 400;
    font-size: 14px;
    color: #D9D9D9;
    line-height: 20px;
    margin-top: 5px;
  }
}
</style>


