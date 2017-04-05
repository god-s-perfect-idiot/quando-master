(function () {
    var self = this["quando_blocks"] = {};
    var PREFIX = 'quando_'; // TODO share with quando_editor
    self.CONFIG = {
        DISPLAY_COLOUR: '#eeee66',
        MEDIA_COLOUR: '#88ee88',
        STYLE_COLOUR: '#eebbee',
        TIME_COLOUR: '#eebbaa',
        LEAP_MOTION_COLOUR: '#bbbbbb',
        BLOCKLY_SATURATION: 0.25, // default for hue only colour - probably not used anymore - see http://colorizer.org/
        BLOCKLY_VALUE: 0.85, // ditto
    };

    var ajax_get = function (url, callback) {
        var xhr = new XMLHttpRequest();
        xhr.onload = function () {
            callback(xhr.responseText);
        };
        xhr.open("GET", url, true);
        xhr.send(null);
    };


    self.addBlocks = function (quando_editor) {

        var STATEMENT = 'STATEMENT';
        var DURATION = 'DURATION';
        var MENU_UNITS = { name: 'Units', title: '', menu: ['Seconds', 'Minutes'] };
        var FREQUENCY = 'FREQUENCY';
        var UNITS_MENU = 'UNITS_MENU';

        quando_editor.defineTime({
            name: 'Check', next: false, previous: false,
            interface: [
                { name: FREQUENCY, title: '', number: 1 },
                {
                    menu: ['Second', 'Minute', 'Hour', 'Day'],
                    name: UNITS_MENU, title: 'times per'
                },
                { statement: STATEMENT }
            ],
            javascript: function (block) {
                var frequency = quando_editor.getNumber(block, FREQUENCY);
                var seconds = 1;
                switch (quando_editor.getMenu(block, UNITS_MENU)) {
                    case 'Minute': seconds = 60;
                        break;
                    case 'Hour': seconds = 60 * 60;
                        break;
                    case 'Day': seconds = 60 * 60 * 24;
                        break;
                };
                var time = seconds / frequency;
                var statement = quando_editor.getStatement(block, STATEMENT);
                var result = "quando.every("
                    + time
                    + ", function() {\n"
                    + statement
                    + "});\n";
                return result;
            }
        });

        var DURATION = 'DURATION';
        var EVERY_BLOCK = 'Every';
        quando_editor.defineTime({
            name: EVERY_BLOCK, next: false, previous: false,
            interface: [
                { name: DURATION, title: '', number: 1 },
                { menu: ['Seconds', 'Minutes', 'Hours'], name: UNITS_MENU, title: '' },
                { statement: STATEMENT }
            ],
            javascript: function (block) {
                var duration = quando_editor.getNumber(block, DURATION);
                switch (quando_editor.getMenu(block, UNITS_MENU)) {
                    case 'Minutes':
                        duration *= 60;
                        break;
                    case 'Hours':
                        duration *= 60 * 60;
                        break;
                };
                block.postfix = '';
                var statement = quando_editor.getStatement(block, STATEMENT);
                var result = "quando.every("
                    + duration
                    + ", function() {\n"
                    + statement + block.postfix
                    + "});\n";
                return result;
            }
        });

        var TIME = 'TIME';
        quando_editor.defineTime({
            name: 'After',
            interface: [
                { name: TIME, title: '', number: 1 },
                { title: 'Seconds' },
                { statement: STATEMENT }
            ],
            javascript: function (block) {
                var time = quando_editor.getNumber(block, TIME); // seconds
                var statement = quando_editor.getStatement(block, STATEMENT);
                var result = "quando.after("
                    + time
                    + ", function() {\n"
                    + statement
                    + "});\n";
                return result;
            }
        });

        var ID_GREETING = "Greeting";
        quando_editor.defineMedia({
            name: 'Show "', title: 'Show Text',
            interface: [{ name: ID_GREETING, title: '"', text: '.type your text here..' }, { title: '"' }],
            javascript: function (block) {
                return 'quando.text("' + quando_editor.getText(block, ID_GREETING) + '");\n';
            }
        });

        var DO_DURATION = 'Do for';
        quando_editor.defineTime({
            name: DO_DURATION,
            interface: [
                { name: DURATION, title: '', number: '1' }, MENU_UNITS,
                { statement: STATEMENT }
            ],
            javascript: function (block) {
                block.postfix = '';
                var seconds = quando_editor.getNumber(block, DURATION) * 1;
                if (quando_editor.getMenu(block, MENU_UNITS.name) === 'Minutes') {
                    seconds *= 60;
                }
                var closing_parent = quando_editor.getParent(block, [EVERY_BLOCK, DO_DURATION, WAIT_ON, FOREVER_BLOCK]);
                var result = '';
                if (quando_editor.getParent(block, [WAIT_ON, FOREVER_BLOCK])) {
                    result += 'inc();\n';
                }
                var statement = quando_editor.getStatement(block, STATEMENT);
                result += "quando.do_duration(" + seconds + ",\n"
                    + "function() {\n"
                    + statement;
                result += block.postfix + "},\n"
                    + "function() {\n";
                if (quando_editor.getParent(block, [WAIT_ON, FOREVER_BLOCK])) {
                    closing_parent.postfix += 'dec();\n';
                }
                closing_parent.postfix += "});\n";
                return result;
            }
        });

        var _getOnContained = function(block, container, contained, otherwise) {
            let result = otherwise;
            if (quando_editor.getParent(block, [container])) {
                result = contained;
            }
            return result;
        }
        var _getStyleOnContained = function(block, container) {
            return 'set' + _getOnContained(block, container, "Display", "Default") + 'Style';
        }

        let COLOUR = 'colour';
        quando_editor.defineStyle({
            name: 'Background', title: 'Background Display Colour',
            interface: [
                { name: COLOUR, title: '', colour: '#ff0000' }
            ],
            javascript: function (block) {
                let method = _getStyleOnContained(block, WHEN_VITRINE_BLOCK);
                let colour = quando_editor.getColour(block, COLOUR);
                return `quando.${method}('#quando_image', 'background-color', '${colour}');\n`;
            }
        });

        var IMAGE = 'Images';
        var FILE_IMAGE = {name: IMAGE, title:'', file:'images'};
        quando_editor.defineMedia({
            name: 'Display', title: 'Show Image',
            interface: [ FILE_IMAGE ],
            javascript: function (block) {
                let method = _getStyleOnContained(block, WHEN_VITRINE_BLOCK);
                let image = quando_editor.getFile(block, IMAGE);
                return `quando.${method}('#quando_image', 'background-image', 'url("/client/images/${image}")');\n`;
            }
        });
        var VIDEO = 'Video';
        var FILE_VIDEO = { name: VIDEO, title: '', file: 'video' };
        quando_editor.defineMedia({
            name: 'Show Video', title: 'Play Video',
            interface: [FILE_VIDEO],
            javascript: function (block) {
                var video_url = quando_editor.getFile(block, VIDEO);
                var result = "quando.video('/client/video/" + video_url + "'";
                if (quando_editor.getParent(block, [WAIT_ON, FOREVER_BLOCK])) {
                    result += ", inc, dec";
                }
                result += ");\n";;
                return result;
            }
        });
        var AUDIO = 'Audio';
        var FILE_AUDIO = {name: AUDIO, title:'', file:'audio'};
        quando_editor.defineMedia({
            name: 'Play Audio',
            interface: [ FILE_AUDIO ],
            javascript: function(block) {
                var _url = quando_editor.getFile(block, AUDIO);
                var result = "quando.audio('/client/audio/" + _url + "'";
                if (quando_editor.getParent(block,[WAIT_ON, FOREVER_BLOCK])) {
                    result += ", inc, dec";
                }
                result += ");\n";;
                return result;
            }
        });
        var CHECK_AUDIO = ' Audio';
        var CHECK_TEXT = ' Text';
        var CHECK_VIDEO = ' Video';
        var CLEAR = 'Clear';
        quando_editor.defineMedia({
            name: CLEAR,
            interface: [
                { name: CHECK_AUDIO, check: false },
                { name: CHECK_TEXT, check: false },
                { name: CHECK_VIDEO, check: false }
            ],
            javascript: function (block) {
                result = "";
                if (quando_editor.getCheck(block, CHECK_AUDIO)) {
                    result += 'quando.clear_audio();\n';
                }
                if (quando_editor.getCheck(block, CHECK_TEXT)) {
                    result += 'quando.text();\n';
                }
                if (quando_editor.getCheck(block, CHECK_VIDEO)) {
                    result += 'quando.clear_video();\n';
                }
                return result;
            }
        });
        var WAIT_ON = 'Wait On';
        quando_editor.defineTime({
            name: WAIT_ON,
            valid_in: [EVERY_BLOCK, DO_DURATION, WAIT_ON, FOREVER_BLOCK],
            interface: [
                { statement: STATEMENT }
            ],
            javascript: function (block) {
                block.postfix = '';
                var statement = quando_editor.getStatement(block, STATEMENT);
                var result = "quando.wait(\n"
                    + " function (inc, dec) {\n"
                    + statement + block.postfix + '},\nfunction() {\n';
                var closing_parent = quando_editor.getParent(block, [EVERY_BLOCK, DO_DURATION, WAIT_ON, FOREVER_BLOCK]);
                closing_parent.postfix += "});\n";
                return result;
            }
        });

        var FOREVER_BLOCK = 'Forever';
        quando_editor.defineTime({
            name: FOREVER_BLOCK, next: false, previous: false,
            interface: [
                { statement: STATEMENT }
            ],
            javascript: function (block) {
                block.postfix = "dec();\n";
                var statement = quando_editor.getStatement(block, STATEMENT);
                var result = "quando.forever(\n"
                    + " function(inc, dec) {\n"
                    + statement + block.postfix + "});";
                return result;

            }
        });
        var LEAP_BLOCK = 'When Hands';
        var HAND_COUNT = 'hand_count';
        quando_editor.defineLeapMotion({
            name: LEAP_BLOCK, next: false, previous: false,
            interface: [
                { name: HAND_COUNT, title: ' = ', number: 1 },
                { statement: STATEMENT }
            ],
            javascript: function (block) {
                var statement = quando_editor.getStatement(block, STATEMENT);
                var result = "quando.hands(" + quando_editor.getNumber(block, HAND_COUNT) + ",\n"
                    + " function() {\n"
                    + statement + "});";
                return result;
            }
        });
        var HANDED_BLOCK = 'When Hand ';
        var HAND_LEFT = 'Left';
        var HAND_RIGHT = 'Right';
        quando_editor.defineLeapMotion({
            name: HANDED_BLOCK, next: false, previous: false,
            interface: [
                { name: HAND_LEFT, check: false },
                { name: HAND_RIGHT, check: false },
                { statement: STATEMENT }
            ],
            javascript: function (block) {
                var statement = quando_editor.getStatement(block, STATEMENT);
                var result = "quando.handed("
                    + quando_editor.getCheck(block, HAND_LEFT) + ",\n"
                    + quando_editor.getCheck(block, HAND_RIGHT) + ",\n"
                    + " function() {\n"
                    + statement + "});";
                return result;
            }
        });

        let DIG_COLOUR = 0;
        let WHEN_VITRINE_BLOCK = 'When Display Case';
        let WHEN_VITRINE_TEXT = 'title';
        quando_editor.defineDisplay({
            name: WHEN_VITRINE_BLOCK, title: 'When Display', next: false, previous: false,
            interface: [{
                name: WHEN_VITRINE_TEXT, title: '', text: 'Title and label',
            },
            { statement: STATEMENT }
            ],
            javascript: (block) => {
                let title = quando_editor.getText(block, WHEN_VITRINE_TEXT);
                var statement = quando_editor.getStatement(block, STATEMENT);
                var result = `quando.vitrine("${block.id}", () => {
quando.title("${title}");
${statement}});
`;
                return result;
            }
        });

        // TODO refactor
        Blockly.mainWorkspace.addChangeListener(function (ev) {
            let workspace = Blockly.Workspace.getById(ev.workspaceId);
            let block = workspace.getBlockById(ev.blockId);
            let topBlocks = Blockly.mainWorkspace.getAllBlocks();
            if (ev.type == Blockly.Events.CHANGE) {
                if (block.type == PREFIX + WHEN_VITRINE_BLOCK) {
                    for (var checkblock of topBlocks) {
                        if ((checkblock.type == PREFIX + LABEL_TO_BLOCK)
                            || (checkblock.type == PREFIX + LABEL_BLOCK)) {
                            let menuid = quando_editor.getMenu(checkblock, LABEL_TO_MENU);
                            if (menuid == block.id) {
                                    quando_editor.setMenuText(checkblock, LABEL_TO_MENU, ev.newValue);
                            }
                        }
                    }
                }
            } else if (ev.type == Blockly.Events.CREATE) {
                for (var checkblock of topBlocks) {
                    if ((checkblock.type == PREFIX + LABEL_TO_BLOCK)
                        || (checkblock.type == PREFIX + LABEL_BLOCK)) {
                        let menuid = quando_editor.getMenu(checkblock, LABEL_TO_MENU);
                        if (menuid == block.id) {
                            quando_editor.setMenuText(checkblock, LABEL_TO_MENU,
                                quando_editor.getText(block, WHEN_VITRINE_TEXT));
                        }
                    }
                }
            } else if (ev.type == Blockly.Events.DELETE) {
                for (var checkblock of topBlocks) {
                    if ((checkblock.type == PREFIX + LABEL_TO_BLOCK)
                        || (checkblock.type == PREFIX + LABEL_BLOCK)) {
                        let menuid = quando_editor.getMenu(checkblock, LABEL_TO_MENU);
                        if (menuid == ev.ids[0]) {
                            quando_editor.resetMenu(checkblock, LABEL_TO_MENU);
                        }
                    }
                }
            }
        });

        // Build the drop down list of Vitrines
        let _label_menu = function () {
            let topBlocks = Blockly.mainWorkspace.getAllBlocks();
            let choices = [['-----', 0]];
            for (var block of topBlocks) {
                if (block.type == PREFIX + WHEN_VITRINE_BLOCK) {
                    let text = quando_editor.getText(block, 'title');
                    choices.push([text, block.id]);
                }
            }
            return choices;
        }
        let LABEL_BLOCK = 'Label';
        let LABEL_TO_MENU = 'to';
        let _label_javascript = function (block) {
            let menuid = quando_editor.getMenu(block, LABEL_TO_MENU);
            // find when block on id, then get it's title
            let whenblock = Blockly.mainWorkspace.getBlockById(menuid);
            let title = quando_editor.getText(whenblock, WHEN_VITRINE_TEXT);
            var result = `quando.addLabel("${menuid}", "${title}");\n`;
            return result;
        }
        let LABEL_TEXT = 'text';
        let LABEL_TO_BLOCK = 'Label to';
        quando_editor.defineDisplay({
            // TODO must be in a vitrine...?
            name: LABEL_TO_BLOCK, title: 'Label',
            interface: [
                {
                    name: LABEL_TO_MENU,
                    menu: _label_menu
                },
            ],
            javascript: _label_javascript,
        });
        let STYLE_BLOCK = 'Style';
        let STYLE_MENU = 'style';
        let DIV_MENU = 'div';
        quando_editor.defineStyle({
            name: STYLE_BLOCK, title: '',
            interface: [
                { menu: ['Title', 'Text', 'Label'], name: DIV_MENU, title: '' },
                {
                    menu: ['Font Colour', 'Background Colour'],
                    name: STYLE_MENU, title: ''
                },
                { name: COLOUR, title: '', colour: '#ff0000' },
            ],
            javascript: (block) => {
                let result ="";
                let method = _getStyleOnContained(block, WHEN_VITRINE_BLOCK);
                let div = quando_editor.getMenu(block, DIV_MENU);
                switch (div) {
                    case 'Title': div = '#quando_title';
                        break;
                    case 'Text': div = '#quando_text';
                        break;
                    case 'Label': div = '.quando_label';
                        break;
                }
                let style = quando_editor.getMenu(block, STYLE_MENU);
                let value = quando_editor.getColour(block, COLOUR);
                if (style == 'Font Colour') {
                    style = 'color';
                } else {
                    style = 'background-color '; // not actually javascript?!
                    // so backgroundColor won't work - has to be CSS interpreted...'
                    let bigint = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(value);
                    let r = parseInt(bigint[1], 16);
                    let g = parseInt(bigint[2], 16);
                    let b = parseInt(bigint[3], 16);
                    value = `rgba(${r}, ${g}, ${b}, 0.6)`;
                    if (div == '.quando_label') { // Need to put in the transition opacity
                        result += `quando.${method}('${div}.focus', '${style}', 'rgba(${r}, ${g}, ${b}, 1)');\n`;
                    }
                }
                result += `quando.${method}('${div}', '${style}', '${value}');\n`;
                return result;
            },
        });
        let FONT_SIZE_BLOCK = 'Font Size';
        let FONT_SIZE = 'font size';
        quando_editor.defineStyle({
            name: FONT_SIZE_BLOCK,
            interface: [
                { menu: ['Title', 'Text', 'Label'], name: DIV_MENU, title: '' },
                { name: FONT_SIZE, title: '', number: 24 }, { title: 'pt' },
            ],
            javascript: (block) => {
                let method = _getStyleOnContained(block, WHEN_VITRINE_BLOCK);
                let div = quando_editor.getMenu(block, DIV_MENU);
                switch (div) {
                    case 'Title': div = '#quando_title';
                        break;
                    case 'Text': div = '#quando_text';
                        break;
                    case 'Label': div = '.quando_label';
                        break;
                }
                let value = quando_editor.getNumber(block, FONT_SIZE) + "pt";
                result = `quando.${method}('${div}', 'font-size', '${value}');\n`;
                return result;
            },
        });

        quando_editor.defineBlock({
            name: 'When ', next: false, previous: false, category: 'extras', colour: '320',
            interface: [
                {
                    menu: ['Diesel', 'Petrol', 'VAT', 'Fuel Tax'],
                    name: UNITS_MENU, title: ''
                },
                { menu: ['Increases', 'Decreases'] },
                { statement: STATEMENT }
            ]
        });

        quando_editor.defineBlock({
            name: 'Text ', category: 'extras', colour: '80',
            interface: [
                { name: 'message', text: 'message' },
                { name: 'to', text: 'mobile' }
            ]
        });

    };
})();