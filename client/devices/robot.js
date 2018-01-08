(function () {
    var quando = this['quando']
    if (!quando) {
      alert('Fatal Error: Robot must be included after quando_browser')
    }

    var self = quando.robot = {}

    var list = null;

    function helper_ConvertAngle(angle) {
        return angle * (Math.PI / 180)
    }

    function helper_TimeoutListening(callback, sr) {
        sr.unsubscribe("NAO_USER");
        callback();
    }

    function conditional(value) {   
        if(value == "disabled" || value == "Stand") 
            return true;
    }

    self.connect = function(robotIp) {
        session = new QiSession(robotIp);       
        session.socket().on('connect', function () {
            console.log('QiSession connected!');
            session.service("ALAutonomousLife").done(function (al) {
                Promise.resolve(al.getState()).then(conditional).then(function (value){
                    if(!value) {
                        session.service("ALTextToSpeech").done(function (tts) {
                            tts.say("Please wait whilst I set up. I only do this once after being turned on, or if you have changed my autonomous life state.");
                          }).fail(function (error) {
                            console.log("An error occurred:", error);
                          });
                        al.setState("disabled");
                    }
                })

                session.service("ALMemory").done(function (ALMemory) {            
                    ALMemory.subscriber("AutonomousLife/State").done(function (sub){
                        sub.signal.connect(function(state){
                            session.service("ALRobotPosture").done(function (rp) {
                                Promise.resolve(rp.getPosture()).then(conditional).then(function (value){
                                    if(!value && state == "disabled") {
                                        rp.goToPosture("Stand", 1.0);
                                    }
                                })                    
                            }).fail(function (error) {
                                console.log("An error occurred:", error);
                            });
                        });
                    });     
                });
            }).fail(function (error) {
                console.log("An error occurred:", error);
            });
            // now you can start using your QiSession
          }).on('disconnect', function () {
            console.log('QiSession disconnected!');
          });
    }

    self.say = function(text, extras) {
        session.service("ALAnimatedSpeech").done(function (as) {
            as.say(text);
          }).fail(function (error) {
            console.log("An error occurred:", error);
          });
    }

    self.moveHead = function(direction, angle) {
        var head;
        switch (direction) {
            case 'Down':
                head = 'HeadPitch';
                break;
            case 'Up':
                head = 'HeadPitch';
                angle = -angle;
                break;  
            case 'Left':
                head = 'HeadYaw';
                break;
            case 'Right':
                head = 'HeadYaw';            
                angle = -angle;            
                break;
        }
        session.service("ALMotion").done(function (motion) {
            newAngle = helper_ConvertAngle(angle);
            motion.setAngles(head, newAngle, 0.5);
          }).fail(function (error) {
            console.log("An error occurred:", error);
          });
    }

    self.moveArm = function(arm, direction, angle) {
        var armJoint;
        var finalAngle;
        switch (arm) {
            case 'Left':
                switch (direction) {
                    case 'Up':
                        armJoint = 'LShoulderPitch';
                        switch (angle) {
                            case 'Halfway':
                                finalAngle = 0;
                                break;
                            case 'Maximum':
                                finalAngle = -90;
                                break;
                        };
                        break;
                    case 'Down':
                        armJoint = 'LShoulderPitch'
                        switch (angle) {
                            case 'Halfway':
                                finalAngle = 0;
                                break;
                            case 'Maximum':
                                finalAngle = 90;
                                break;
                        };
                        break; 
                    case 'Left':
                        armJoint = 'LShoulderRoll'
                        switch (angle) {
                            case 'Halfway':
                                finalAngle = 30;
                                break;
                            case 'Maximum':
                                finalAngle = 60;
                                break;
                        };
                        break;
                    case 'Right':
                        armJoint = 'LShoulderRoll'
                        switch (angle) {
                            case 'Halfway':
                                finalAngle = 0;
                                break;
                            case 'Maximum':
                                finalAngle = -18;
                                break;
                            };              
                        break;  
                }
                break;
            case 'Right':
            switch (direction) {
                case 'Up':
                    armJoint = 'RShoulderPitch';
                    switch (angle) {
                        case 'Halfway':
                            finalAngle = 0;
                            break;
                        case 'Maximum':
                            finalAngle = -90;
                            break;
                    };
                    break;
                case 'Down':
                    armJoint = 'RShoulderPitch'
                    switch (angle) {
                        case 'Halfway':
                            finalAngle = 0;
                            break;
                        case 'Maximum':
                            finalAngle = 90;
                            break;
                    };
                    break; 
                case 'Left':
                    armJoint = 'RShoulderRoll'
                    switch (angle) {
                        case 'Halfway':
                            finalAngle = 0;
                            break;
                        case 'Maximum':
                            finalAngle = 18;
                            break;
                    };
                    break;
                case 'Right':
                armJoint = 'RShoulderRoll'
                    switch (angle) {
                        case 'Halfway':
                            finalAngle = -30;
                            break;
                        case 'Maximum':
                            finalAngle = -60;
                            break;
                        };              
                    break;  
                }
                break;                  
            }
        session.service("ALMotion").done(function (motion) {
            newAngle = helper_ConvertAngle(finalAngle);
            motion.setAngles(armJoint, newAngle, 0.5);
          }).fail(function (error) {
            console.log("An error occurred:", error);
          });
    }

    self.motionHand = function(hand, open) {
        session.service("ALMotion").done(function (motion) {
            if(open=='Open') {
                motion.openHand(hand);
            } else {
                motion.closeHand(hand);
            }
          }).fail(function (error) {
            console.log("An error occurred:", error);
          });
    }

    self.personPerception = function(callback) {
        session.service("ALMemory").done(function (ALMemory) {            
            ALMemory.subscriber("ALBasicAwareness/HumanTracked").done(function (sub){
                sub.signal.connect(function(state){
                    console.log("Found You!");
                    callback();
                });
            });     
        });
    }

    self.changeAutonomousLife = function(state) {
        session.service("ALAutonomousLife").done(function (al) {
            al.setState(state);
          }).fail(function (error) {
            console.log("An error occurred:", error);
          });
    }

    self.setVocabulary = function(vocab) {
        list = vocab.split(",");
    }

    self.listenForWords = function(callback, callback2, timeout) {
        session.service("ALSpeechRecognition").done(function (sr) {
            sr.setVocabulary(list, true);
            sr.pause(false);
            if (sr != null && list != null) {
                sr.subscribe("NAO_USER");
            } else {
                alert("You haven't set any vocabulary!");
            } 

            session.service("ALMemory").done(function (ALMemory) {            
                ALMemory.subscriber("WordRecognized").done(function (sub){
                    sub.signal.connect(function(value){
                        console.log("I recognise that word!");
                        console.log(value);  
                        callback();                  
                    });
                });
            });
            setTimeout(function() { helper_TimeoutListening(callback2, sr); }, timeout*1000);                    
          }).fail(function (error) {
            console.log("An error occurred:", error);
          }); 
    }

    self.stopListening = function(callback) {
        session.service("ALSpeechRecognition").done(function (sr) {
            sr.unsubscribe("NAO_USER");
        });
    }
})()
