var toml = require("toml");
var S = require("string");

var CONTENT_PATH_PREFIX = "content",
    SCSS_DIR = "themes/cleverlee/scss",
    SCSS_PATH =SCSS_DIR + "/*.scss",
    CSS_DIR = "themes/cleverlee/static/css",
    BULMA_DIR = "node_modules/bulma",
    SEARCH_JSON_FILE = "content/search.json";

// 载入模块
var Segment = require('segment');
// 创建实例
var segment = new Segment();
// 使用默认的识别模块及字典，载入字典文件需要1秒，仅初始化时执行一次即可
segment.useDefault();

module.exports = function(grunt) {
    // 使用严格模式
    'use strict';

    // 这里定义我们需要的任务
    grunt.initConfig({

        // 设置任务，删除文件夹
        clean: {
            dist: CSS_DIR
        },


        // 通过sass编译成css文件
        sass: {
            dist: {
                options: {
                    loadPath: [BULMA_DIR],
                    style: 'compressed'
                },
                files: [{
                    expand: true,
                    cwd: SCSS_DIR,
                    src: ['*.scss'],
                    dest: CSS_DIR,
                    ext: '.css'
                }]
            }
        },

        // 检测改变，自动跑sass任务
        watch: {
            scripts: {
                files: [SCSS_PATH, CONTENT_PATH_PREFIX + '/**/*.md'],
                tasks: ['sass', 'lunr-search'],
                options: {
                    spawn: false
                }
            }
        }
    });

    // 一定要引用着3个模块
    grunt.loadNpmTasks('grunt-contrib-clean');
    grunt.loadNpmTasks('grunt-contrib-sass');
    grunt.loadNpmTasks('grunt-contrib-watch');
    // 把需要跑的任务注册到default这里每次运行grunt的时候先删除dist，然后重新编译，最后监测文件夹的情况。
    grunt.registerTask('default', ['clean:dist', 'sass:dist', 'watch:scripts']);

    ///////////////////////search index

    grunt.registerTask("lunr-search", function() {

        grunt.log.writeln("Build pages index");

        var indexPages = function() {
            var pagesIndex = [];
            grunt.file.recurse(CONTENT_PATH_PREFIX, function(abspath, rootdir, subdir, filename) {
                grunt.verbose.writeln("Parse file:",abspath);
                var result = processFile(abspath, filename);
                if (result){
                    pagesIndex.push(result);
                }
            });

            return pagesIndex;
        };

        var processFile = function(abspath, filename) {
            var pageIndex;

            if (S(filename).endsWith(".html")) {
                pageIndex = processHTMLFile(abspath, filename);
            } else {
                pageIndex = processMDFile(abspath, filename);
            }

            return pageIndex;
        };

        var processHTMLFile = function(abspath, filename) {
            var content = grunt.file.read(abspath);
            var pageName = S(filename).chompRight(".html").s;
            var href = S(abspath)
                .chompLeft(CONTENT_PATH_PREFIX).s;
            return {
                title: pageName,
                href: href.replace(/ /g, '-').replace(/（/g, '').replace(/）/g, '')
                    .replace(/\(/g, '').replace(/\)/g, ''),
                content: segment.doSegment(pageName + S(content[2]).trim().stripTags(), {
                    simple: true,
                    stripPunctuation: true
                })
            };
        };

        var processMDFile = function(abspath, filename) {
            var content = grunt.file.read(abspath);
            var pageIndex;
            // First separate the Front Matter from the content and parse it
            content = content.split("+++");

            var frontMatter;
            try {
                frontMatter = toml.parse(content[1].trim());
            } catch (e) {
                console.error(e.message);
            }

            if (frontMatter == undefined) {
                return;
            }

            var href = S(abspath).chompLeft(CONTENT_PATH_PREFIX).chompRight(".md").s;
            // href for index.md files stops at the folder name
            if (filename === "index.md") {
                href = S(abspath).chompLeft(CONTENT_PATH_PREFIX).chompRight(filename).s;
            }

            // Build Lunr index for this page
            pageIndex = {
                title: frontMatter.title,
                tags: frontMatter.tags,
                href: href.replace(/ /g, '-').replace(/（/g, '').replace(/）/g, '')
                    .replace(/\(/g, '').replace(/\)/g, ''),
                content: segment.doSegment(frontMatter.title + S(content[2]).trim().stripTags(), {
                    simple: true,
                    stripPunctuation: true
                })
            };

            return pageIndex;
        };

        grunt.file.write(SEARCH_JSON_FILE, JSON.stringify(indexPages()));
        grunt.log.ok("Index built");
    });
};

