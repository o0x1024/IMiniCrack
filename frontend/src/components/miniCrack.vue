
<template>
  <!-- <a-row>
    <a-col :span="24">
      <a-upload-dragger v-model:fileList="fileList" name="file" :multiple="true"
        action="https://www.mocky.io/v2/5cc8019d300000980a055e76" @change="handleChange" @drop="handleDrop">
        <p class="ant-upload-drag-icon">
          <inbox-outlined></inbox-outlined>
        </p>
        <p class="ant-upload-text">Click or drag file to this area to upload</p>
        <p class="ant-upload-hint">
          Support for a single or bulk upload. Strictly prohibit from uploading company data or other
          band files
        </p>
      </a-upload-dragger>
    </a-col>
  </a-row> -->

  <div style="padding: 10px;">
    <a-row gutter="10">
      <a-col>
        <p style="padding: 5px;">小程序包路径：</p>
      </a-col>
      <a-col>
        <a-input v-model:value="wxPath" style="width: 500px;"
          placeholder="C:\Users\test\Documents\WeChat Files\Applet" />
      </a-col>
      <c-col :span="2">
        <a-button @click="btnChange">
          <template #icon>
            <UploadOutlined />
          </template>
          选择文件
        </a-button>
      </c-col>
    </a-row>

    <a-row gutter="10">
      <a-col>
        <p style="padding: 5px;">WXID：&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p>
      </a-col>
      <a-col>
        <a-input v-model:value="wxId" style="width: 300px;" placeholder="wxId" />
      </a-col>
    </a-row>


    <a-row :gutter="10">
      <a-col>
        <p style="padding: 5px;">导出路径：&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p>
      </a-col>
      <a-col style="padding: 5px;">
        <p class="outPath">{{ outPath }}</p>
      </a-col>
      <a-col>
        <a-button @click="btnSelectOutPath">
          <template #icon>
            <UploadOutlined />
          </template>
          选择目录
        </a-button>
      </a-col>
      <a-col>
        <a-button type="primary" @click="btnOpenDecDir">打开解密目录</a-button>
      </a-col>
    </a-row>



    <a-row>
      <a-col>
        <a-button type="primary" @click="btnDecExport">解密导出</a-button>
      </a-col>

    </a-row>


    <a-row style="margin-top: 10px;">
      <a-col :span="24">
        <!-- :disabled="true" -->
        <a-textarea v-model:value="logger" placeholder="日志" :rows="23" />
      </a-col>
    </a-row>
  </div>


</template>


<script lang="ts">
import { computed, onMounted, reactive } from 'vue'
import { OpenFile, OpenDir,OpenDecDir } from "../../wailsjs/go/main/App"
import { Unpack } from "../../wailsjs/go/crack/Crack"
import { EventsOn } from "../../wailsjs/runtime/runtime"
import { InboxOutlined, UploadOutlined } from '@ant-design/icons-vue';
import { message } from 'ant-design-vue';
import { defineComponent, ref } from 'vue';



export default defineComponent({
  components: {
    InboxOutlined,
    UploadOutlined
  },
  setup() {


    onMounted(() => {
      EventsOn("log", (data: any) => {
        if (data) {
          logger.value += data+"\n"
          // console.log(data)
        }
      })

    })


    let wxId = ref('')
    let wxPath = ref('')
    let outPath = ref('C:\\')
    let logger = ref('')


    wxId = computed(() => {
      let pathSlice = wxPath.value.split("\\")
      if (pathSlice.length <= 0) {
        pathSlice = wxPath.value.split("/")
      }

      for (let i in pathSlice) {

        if (pathSlice[i].indexOf("wx") != -1) {
          return pathSlice[i]
        }
      }
      return ""
    })

    const btnSelectOutPath = (e: any) => {
      OpenDir().then((result) => {
        outPath.value = result
      })
    }

    const btnChange = () => {
      OpenFile().then((result) => {
        if (result) {
          wxPath.value = result
        }
      })
    }

    const btnDecExport = () => {
      Unpack(wxPath.value, wxId.value, outPath.value).then((result) => {
        if (result) {
          message.info(result)
        }
      })
    }

    const btnOpenDecDir = () => {
      OpenDecDir(outPath.value)
    }

    return {
      logger,
      wxId,
      outPath,
      wxPath,
      btnChange,
      btnDecExport,
      btnOpenDecDir,
      btnSelectOutPath,
    };
  },
});

</script>


<style scoped>

</style>
