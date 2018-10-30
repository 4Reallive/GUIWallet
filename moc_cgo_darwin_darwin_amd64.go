// +build !ios

package main

/*
#cgo CFLAGS: -pipe -O2 -arch x86_64 -isysroot /Library/Developer/CommandLineTools/SDKs/MacOSX10.14.sdk -mmacosx-version-min=10.11 -Wall -W -fPIC -DQT_NO_DEBUG -DQT_MULTIMEDIA_LIB -DQT_DESIGNER_LIB -DQT_UIPLUGIN_LIB -DQT_WIDGETS_LIB -DQT_QUICK_LIB -DQT_GUI_LIB -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_DBUS_LIB -DQT_XML_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -stdlib=libc++ -O2 -std=gnu++11 -arch x86_64 -isysroot /Library/Developer/CommandLineTools/SDKs/MacOSX10.14.sdk -mmacosx-version-min=10.11 -Wall -W -fPIC -DQT_NO_DEBUG -DQT_MULTIMEDIA_LIB -DQT_DESIGNER_LIB -DQT_UIPLUGIN_LIB -DQT_WIDGETS_LIB -DQT_QUICK_LIB -DQT_GUI_LIB -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_DBUS_LIB -DQT_XML_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../src -I. -I../../../Qt5.11.1/5.11.1/clang_64/lib/QtMultimedia.framework/Headers -I../../../Qt5.11.1/5.11.1/clang_64/lib/QtDesigner.framework/Headers -I../../../Qt5.11.1/5.11.1/clang_64/lib/QtUiPlugin.framework/Headers -I../../../Qt5.11.1/5.11.1/clang_64/lib/QtWidgets.framework/Headers -I../../../Qt5.11.1/5.11.1/clang_64/lib/QtQuick.framework/Headers -I../../../Qt5.11.1/5.11.1/clang_64/lib/QtGui.framework/Headers -I../../../Qt5.11.1/5.11.1/clang_64/lib/QtQml.framework/Headers -I../../../Qt5.11.1/5.11.1/clang_64/lib/QtNetwork.framework/Headers -I../../../Qt5.11.1/5.11.1/clang_64/lib/QtDBus.framework/Headers -I../../../Qt5.11.1/5.11.1/clang_64/lib/QtXml.framework/Headers -I../../../Qt5.11.1/5.11.1/clang_64/lib/QtCore.framework/Headers -I. -I/Library/Developer/CommandLineTools/SDKs/MacOSX10.14.sdk/System/Library/Frameworks/OpenGL.framework/Headers -I/Library/Developer/CommandLineTools/SDKs/MacOSX10.14.sdk/System/Library/Frameworks/AGL.framework/Headers -I../../../Qt5.11.1/5.11.1/clang_64/mkspecs/macx-clang -F/Users/jarrod/Qt5.11.1/5.11.1/clang_64/lib
#cgo LDFLAGS: -stdlib=libc++ -headerpad_max_install_names -arch x86_64 -Wl,-syslibroot,/Library/Developer/CommandLineTools/SDKs/MacOSX10.14.sdk -mmacosx-version-min=10.11 -Wl,-rpath,@executable_path/Frameworks -Wl,-rpath,/Users/jarrod/Qt5.11.1/5.11.1/clang_64/lib
#cgo LDFLAGS:  -F/Users/jarrod/Qt5.11.1/5.11.1/clang_64/lib -framework QtMultimedia -framework QtNetwork -framework QtCore -framework DiskArbitration -framework IOKit -framework QtGui -framework QtDesigner -framework QtWidgets -framework QtXml -framework QtQuick -framework QtQml -framework QtDBus -framework OpenGL -framework AGL
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
