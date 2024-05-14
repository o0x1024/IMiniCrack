<template>
    <Codemirror class="codemirror" v-model:value="code" :options="cmOptions" border placeholder="test placeholder"
        :height="660" @change="change" />
</template>

<script lang="ts">
import Codemirror from "codemirror-editor-vue3";
import { defineComponent, ref, computed, watch, onMounted } from 'vue';
import "codemirror/mode/javascript/javascript.js";
import "codemirror/theme/dracula.css";


window.addEventListener('touchstart', function (event) {
    // some logic
    event.preventDefault(); // <-- that should not be used in passive
    // some other magic
});

window.addEventListener('touchmove', function (event) {
    // some logic
    event.preventDefault(); // <-- that should not be used in passive
    // some other magic
});

export default defineComponent({
    props: {
        Data: {
            type: String
        },
        LineNum: {
            type: String
        }
    },
    name: "MyCodeMirror",
    components: { Codemirror },
    setup(props) {
        const code = computed({
            get() {
                return props.Data
            },
            set(val) {

            }
        })

        setInterval(() => {
            code.value += "test \n";
        }, 300);


        watch(() => code, (newcode: any) => {
            console.log(newcode)

        })

        onMounted(() => {
            console.log("mouasdnted")
        })

        const change = (val: any, instance: any) => {

            console.log(props.LineNum)
            var line = instance.getLineHandle(props.LineNum);
            var lineStart = { line: Number(props.LineNum)-1, ch: 0 };
            var lineEnd = { line: Number(props.LineNum), ch: 0 };
            instance.markText(lineStart, lineEnd, { className: "code-highlight" });
            instance.scrollIntoView( Number(props.LineNum)+20);
        }

        return {
            code,
            change,
            cmOptions: {
                mode: "text/javascript", // Language mode
                theme: "dracula", // Theme
                lineNumbers: true, // Show line number
                smartIndent: true, // Smart indent
                indentUnit: 4, // The smart indent unit is 2 spaces in length
                foldGutter: true, // Code folding
                styleActiveLine: true, // Display the style of the selected row
            },
        };
    },
});
</script>

<style>
.codemirror {
    touch-action: none;
}


.code-highlight {
    background-color: rgb(80, 78, 78);
  }
</style>