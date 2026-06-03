<template>
    <template v-if="maxNumber <= 1">
      <div class="single-upload-btn" @click="showModal = true" v-if="fileList.length <= 0">点击上传</div>
      <div class="single-success-image" v-else>
        <img class="single-success-image-show" :src="fileList[0]" @error="errorImg($event)" alt="" />
        <div class="single-success-image-action">
          <n-button
            text
            style="margin-right: 10px"
            :style="buttonCSS"
            @click="handlePreview(fileList[0])"
          >
            <n-icon size="18">
              <EyeOutline />
            </n-icon>
          </n-button>
          <n-button
            text
            :style="buttonCSS"
            @click="handleRemove(0)"
          >
            <n-icon size="18">
              <TrashBinOutline />
            </n-icon>
          </n-button>
        </div>
      </div>
    </template>
    <template v-else>
      <div class="multiple-image-div">
        <div class="multiple-success-image" v-for="(item, index) in fileList" :key="`img_${index}`">
          <img class="multiple-success-image-show" :src="item" @error="errorImg($event)" alt="" />
          <div class="multiple-success-image-action">
            <n-button
              text
              style="margin-right: 10px"
              :style="buttonCSS"
              @click="handlePreview(item)"
            >
              <n-icon size="18">
                <EyeOutline />
              </n-icon>
            </n-button>
            <n-button
              text
              :style="buttonCSS"
              @click="handleRemove(index)"
            >
              <n-icon size="18">
                <TrashBinOutline />
              </n-icon>
            </n-button>
          </div>
        </div>
        <div class="multiple-upload-btn" @click="showModal = true">
          <n-icon size="24">
            <AddOutline />
          </n-icon>
        </div>
      </div>
    </template>

  <n-modal
    v-model:show="showModal"
    :on-after-leave="handleCancel"
    :style="{
      width: dialogWidth,
    }"
  >
    <n-card title="文件选择">
      <template #header-extra>
        <n-space>
          <n-button @click="handleUpload" ghost>
            <template #icon>
              <n-icon :component="UploadOutlined" />
            </template>
            上传文件
          </n-button>
        </n-space>
      </template>
      <n-card style="overflow: auto" content-style="padding: 0;">
        <Chooser
          ref="chooserRef"
          :file-type="fileType"
          :maxNumber="maxNumber"
          :fileList="fileList"
          @saveChange="saveChange"
          @saveChange2="saveChange2"
          @changeKind="changeKind"
        />
      </n-card>
      <template #footer>
        <n-space justify="end">
          <n-button @click="handleCancel"> 取消 </n-button>
          <n-button type="primary" @click="handleSelectFile"> 确定 </n-button>
        </n-space>
      </template>
    </n-card>
  </n-modal>

  <FileUpload
    ref="fileUploadRef"
    :width="dialogWidth"
    :finish-call="handleFinishCall"
    :max-upload="20"
    :file-kind="fileKind"
  />

  <Preview ref="previewRef" />
</template>

<script lang="ts" setup>
  import { NButton, NSpace, NCard, NModal, NIcon, useDialog } from 'naive-ui';
  import { cloneDeep } from 'lodash-es';
  import FileUpload from '@/components/FileChooser/src/Upload.vue';
  import Chooser from '@/components/FileChooser/src/Chooser.vue';
  import Preview from '@/components/FileChooser/src/Preview.vue';
  import { computed, onMounted, ref, watch } from 'vue';
  import { adaModalWidth, errorImg } from '@/utils/hotgo';
  import { getFileExt } from '@/utils/urlUtils';
  import {
    UploadOutlined,
    PlusOutlined,
    CloudDownloadOutlined,
    DeleteOutlined,
    EyeOutlined,
  } from '@vicons/antd';
  import {
    EyeOutline,
    TrashBinOutline,
    AddOutline
  } from '@vicons/ionicons5'
  import { Attachment, FileType, getFileType } from '@/components/FileChooser/src/model';
  import { isArrayString, isString } from '@/utils/is';

  export interface Props {
    value: string | string[] | null;
    maxNumber?: number;
    fileType?: FileType;
    width?: number;
    height?: number;
    nolist?: boolean;
  }

  const props = withDefaults(defineProps<Props>(), {
    value: '',
    maxNumber: 1,
    fileType: 'default',
    width: 100,
    height: 100,
  });

  const emit = defineEmits(['update:value', 'change']);
  const fileKind = ref('');
  const fileUploadRef = ref();
  const dialog = useDialog();
  const showModal = ref(false);
  const chooserRef = ref();
  const previewRef = ref();
  const fileList = ref<string[]>([]);
  const fileListall = ref<string[]>([]);
  const dialogWidth = computed(() => {
    return adaModalWidth(1080);
  });

  const getCSSProperties = computed(() => {
    return {
      width: `${props.width}px`,
      height: `${props.height}px`,
    };
  });

  const fileAvatarCSS = computed(() => {
    return {
      '--n-merged-size': `var(--n-avatar-size-override, ${props.width * 0.8}px)`,
      '--n-font-size': `18px`,
    };
  });

  const buttonCSS = computed(() => {
    return {
      '--n-text-color': 'rgb(51, 54, 57)',
      '--n-text-color-hover': 'rgb(51, 54, 57)',
      '--n-text-color-pressed': 'rgb(51, 54, 57)',
      '--n-text-color-focus': 'rgb(51, 54, 57)',
    }
  })

  const buttonText = computed(() => {
    return getFileType(props.fileType);
  });

  // 预览
  function handlePreview(url: string) {
    previewRef.value.openPreview(url);
  }

  // 下载
  function handleDownload(url: string) {
    window.open(url);
  }

  // 删除
  function handleRemove(index: number) {
    dialog.info({
      title: '提示',
      content: '你确定要删除吗？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        fileList.value.splice(index, 1);
        if (props.maxNumber === 1) {
          emit('update:value', '');
          emit('change', '');
        } else {
          emit('update:value', fileList.value);
          emit('change', fileList.value);
        }
      },
      onNegativeClick: () => {},
    });
  }

  function handleSelectFile() {
    showModal.value = false;
    if (props.maxNumber === 1) {
      emit('update:value', fileList.value.length > 0 ? fileList.value[0] : '');
      emit('change', {
        value: fileList.value.length > 0 ? fileList.value[0] : '',
        valueall: fileUploadRef.value.length > 0 ? fileListall.value[0] : '',
      });
    } else {
      emit('update:value', fileList.value);
      emit('change', { value: fileList.value, allvalue: fileListall.value });
    }
  }

  function handleCancel() {
    showModal.value = false;
    loadImage();
  }

  function handleUpload() {
    fileUploadRef.value.openModal();
  }

  function handleFinishCall(result: Attachment, success: boolean) {
    if (success) {
      chooserRef.value.reloadTable();
    }
  }

  function saveChange(list: string[]) {
    fileList.value = list;
  }
  function saveChange2(list: string[]) {
    fileListall.value = list;
  }
  function changeKind(kind){
    fileKind.value = kind
  }
  function loadImage() {
    const value = cloneDeep(props.value);
    if (props.maxNumber === 1) {
      fileList.value = [];
      if (value !== '') {
        if (!isString(value)) {
          console.warn(
            'When the file picker is currently in single-file mode, but the passed value is not of type string, there may be potential issues.'
          );
        }
        fileList.value.push(value as string);
      }
    } else {
      if (!isArrayString(value)) {
        console.warn(
          'When the file picker is currently in multiple-file mode, but the passed value is not of type string array, there may be potential issues.'
        );
      }
      if (!value) {
        fileList.value = [];
      } else {
        fileList.value = value as string[];
      }
    }
  }

  watch(
    () => props.value,
    () => {
      loadImage();
    },
    {
      immediate: true,
      deep: true,
    }
  );

  onMounted(async () => {
    loadImage();
  });
</script>

<style lang="less">
  .n-upload {
    width: auto; /**  居中 */
  }

  .upload {
    width: 100%;
    overflow: hidden;

    &-card {
      width: auto;
      height: auto;
      display: flex;
      flex-wrap: wrap;
      align-items: center;

      &-item {
        margin: 0 8px 8px 0;
        position: relative;
        padding: 8px;
        border: 1px solid #d9d9d9;
        border-radius: 2px;
        display: flex;
        justify-content: center;
        flex-direction: column;
        align-items: center;

        &:hover {
          background: 0 0;

          .upload-card-item-info::before {
            opacity: 1;
          }

          &-info::before {
            opacity: 1;
          }
        }

        &-info {
          position: relative;
          height: 100%;
          padding: 0;
          overflow: hidden;

          &:hover {
            .img-box-actions {
              opacity: 1;
            }
          }

          &::before {
            position: absolute;
            z-index: 1;
            width: 100%;
            height: 100%;
            background-color: rgba(0, 0, 0, 0.5);
            opacity: 0;
            transition: all 0.3s;
            content: ' ';
          }

          .img-box {
            position: relative;
            //padding: 8px;
            //border: 1px solid #d9d9d9;
            border-radius: 2px;
          }

          .img-box-actions {
            position: absolute;
            top: 50%;
            left: 50%;
            z-index: 10;
            white-space: nowrap;
            transform: translate(-50%, -50%);
            opacity: 0;
            transition: all 0.3s;
            display: flex;
            align-items: center;
            justify-content: space-between;

            &:hover {
              background: 0 0;
            }

            .action-icon {
              color: rgba(255, 255, 255, 0.85);

              &:hover {
                cursor: pointer;
                color: #fff;
              }
            }
          }
        }
      }

      &-item-select-picture {
        border: 1px dashed #d9d9d9;
        border-radius: 2px;
        cursor: pointer;
        background: #fafafa;
        color: #666;

        .upload-title {
          color: #666;
        }
      }
    }
  }

  .single-upload-btn{
    width: 95px;
    height: 95px;
    border-radius: 2px;
    background: #FAFAFC;
    border: 1px dashed #E0E0E6;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: 400;
    font-size: 14px;
    color: #3D3D3D;
    line-height: 20px;
    cursor: pointer;
    &:hover{
      border-color: #053dc8;
    }
  }

  .single-success-image{
    position: relative;
    width: 95px;
    height: 95px;
    border-radius: 2px;
    border: 1px solid #E0E0E6;
    .single-success-image-show{
      width: 100%;
      height: 100%;
      border-radius: 2px;
    }
  }
  .single-success-image-action{
    z-index: 2;
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    background: rgba(224,224,230,0);
    opacity: 0;
    &:hover{
      background: rgba(224,224,230,0.7);
      opacity: 1;
    }
  }

  .multiple-image-div{
    overflow: hidden;
    .multiple-success-image{
      float: left;
      position: relative;
      width: 95px;
      height: 95px;
      border-radius: 2px;
      border: 1px solid #E0E0E6;
      margin-right: 10px;
      margin-bottom: 10px;
      .multiple-success-image-show{
        width: 100%;
        height: 100%;
        border-radius: 2px;
      }
      .multiple-success-image-action{
        z-index: 2;
        position: absolute;
        width: 100%;
        height: 100%;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        display: flex;
        align-items: center;
        justify-content: center;
        background: rgba(224,224,230,0);
        opacity: 0;
        &:hover{
          background: rgba(224,224,230,0.7);
          opacity: 1;
        }
      }
    }

    .multiple-upload-btn{
      width: 95px;
      height: 95px;
      border-radius: 2px;
      background: #FAFAFC;
      border: 1px dashed #E0E0E6;
      display: flex;
      align-items: center;
      justify-content: center;
      font-weight: 400;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      cursor: pointer;
      &:hover{
        border-color: #053dc8;
      }
    }
  }

</style>
