package main

import (
	"github.com/oderwat/timg/must"
	"path/filepath"
	"strings"
)

type extType int

const (
	extSkipped extType = iota
	extUnknown
	extOK
)

func extensionType(filename string) extType {
	imgExt := map[string]struct{}{".jpg": {}, ".jpeg": {}, ".png": {}, ".gif": {}, ".bmp": {}, ".webp": {}}
	skipExt := map[string]struct{}{"": {}, ".mp4": {}, ".mpg": {}, ".mpeg": {}, ".wmv": {}, ".3gp": {}, ".mov": {},
		".asf": {}, ".flv": {}, ".m4v": {}, ".avi": {}, ".mkv": {}, ".webm": {}, ".sh": {}, ".db": {}, ".part": {},
		".dmg": {}, ".rar": {}, ".zip": {}, ".apk": {}, ".html": {}, ".txt": {}, ".stl": {}, ".f3d": {}, ".factory": {},
		".xml": {}, ".bat": {}, ".strings": {}, ".plist": {}, ".icns": {}, ".reg": {}, ".fon": {}, ".dll": {},
		".so": {}, ".ttf": {}, ".inf": {}, ".exe": {}, ".com": {}, ".dylib": {}, ".sc": {}, ".wad": {}, ".itc": {},
		".m4p": {}, ".py": {}, ".pyc": {}, ".ctype": {}, ".fish": {}, ".db-shm": {}, ".db-wal": {}, ".nib": {},
		".rst": {}, ".pth": {}, ".h": {}, ".c": {}, ".csh": {}, ".pst": {}, ".json": {}, ".js": {}, ".sbstore": {},
		".vlpset": {}, ".metadata": {}, ".ftl": {}, ".bin": {}, ".little": {}, ".lz4": {}, ".sqlite3": {},
		".sqlite3-shm": {}, ".sqlite3-wal": {}, ".pbd": {}, ".otf": {}, ".icondatapack": {}, ".iconconfigpack": {},
		".styl": {}, ".c3b": {}, ".c3h": {}, ".crate": {}, ".hs": {}, ".cleaned": {}, ".iconmappack": {},
		".sample": {}, ".package": {}, ".package-cache": {}, ".idx": {}, ".pack": {}, ".bif": {}, ".0": {},
		".ini": {}, ".lic": {}, ".dat": {}, ".ast": {}, ".unlinked2": {}, ".declarations": {}, ".linked_bundle": {},
		".instrumentation": {}, ".pub-package-details-cache": {}, ".log": {}, ".yaml": {}, ".stfolder": {},
		".gypi": {}, ".java": {}, ".php": {}, ".phtml": {}, ".declarations_content": {},
		".resolved": {}, ".yml": {}, ".bashrc": {}, ".profile": {}, ".cache-7": {}, ".lock": {}, ".cer": {},
		".crt": {}, ".pem": {}, ".key": {}, ".groovy": {}, ".conf": {}, ".scala": {}, ".gradle": {}, ".css": {},
		".kts": {}, ".map": {}, ".ts": {}, ".msi": {}, ".url": {}, ".etag": {}, ".pub": {}, ".md": {},
		".github": {}, ".htaccess": {}, ".svg": {}, ".kt": {}, ".rb": {}, ".ignore-me": {}, ".tmp": {}, ".pom": {},
		".bash_profile": {}, ".srl": {}, ".htm": {}, ".gitattributes": {}, ".go": {}, ".mod": {}, ".work": {},
		".swift": {}, ".properties": {}, ".jar": {}, ".md5": {}, ".sha1": {}, ".gpg": {}, ".plugin": {},
		".cnf": {}, ".img": {}, ".state": {}, ".npmignore": {}, ".mjs": {}, ".bnf": {}, ".licenses": {},
		".blocks": {}, ".session": {}, ".history": {}, ".historynew": {}, ".avd": {}, ".cpp": {}, ".routes": {},
		".coffee": {}, ".lockfile": {}, ".pid": {}, ".d": {}, ".jsp": {}, ".hpp": {}, ".lib": {},
		".ps": {}, ".ronn": {}, ".gemspec": {}, ".gitignore_global": {}, ".example": {}, ".rc": {}, ".a": {},
		".s": {}, ".mm": {}, ".idl": {}, ".module": {}, ".processor": {}, ".vst": {}, ".back": {}, ".woff": {},
		".app": {}, ".config": {}, ".rights": {}, ".jfc": {}, ".policy": {}, ".notice": {}, ".cur": {}, ".tt": {},
		".pump": {}, ".lck": {}, ".ok": {}, ".remote": {}, ".dtd": {}, ".desktop": {}, ".pf": {}, ".data": {},
		".src": {}, ".dir": {}, ".net": {}, ".org": {}, ".document": {}, ".pro": {}, ".pluginserviceregistry": {},
		".sqlite": {}, ".cfg": {}, ".bfc": {}, ".template": {}, ".access": {}, ".gnu": {}, ".npz": {}, ".cache": {},
		".g": {}, ".bundler": {}, ".standalone": {}, ".xsl": {}, ".extensionbundle": {}, ".processors": {}, ".csv": {},
		".m": {}, ".vcxproj": {}, ".filters": {}, ".bsd": {}, ".get": {}, ".ja": {}, ".certs": {}, ".libraries": {},
		".security": {}, ".psd": {}, ".war": {}, ".pl": {}, ".markdown": {}, ".cjs": {}, ".editorconfig": {},
		".gitkeep": {}, ".def": {}, ".jst": {}, ".less": {}, ".priv": {}, ".jshintrc": {}, ".jsx": {}, ".tsx": {},
		".svelte": {}, ".vue": {}, ".xz": {}, ".ls": {}, ".sass": {}, ".scss": {}, ".flow": {}, ".opts": {},
		".flags": {}, ".nycrc": {}, ".node": {}, ".extensionmodule": {}, ".lcl": {}, ".eslintrc": {}, ".ipynb": {},
		".po": {}, ".mo": {}, ".mplstyle": {}, ".dist-info": {}, ".ics": {}, ".gpx": {}, "stl": {}, ".afdesign": {},
		".mix": {}, ".eml": {}, ".calendar": {}, ".ca": {}, ".csr": {}, ".icsalarm": {}, ".gcode": {}, ".tpl": {},
		".de": {}, ".z": {}, ".fdf": {}, ".xls": {}, ".xlsx": {}, ".cs": {}, ".docx": {}, ".wdsl": {}, ".xsd": {},
		".dist": {}, ".xlf": {}, ".adsklib": {}, ".exr": {}, ".out": {}, ".b": {}, ".min": {}, ".cc": {}, ".ogg": {},
		".pak": {}, ".obj": {}, ".prt": {}, ".patch": {}, ".diff": {}, ".sav": {}, ".cube_shaperlut": {}, ".ldb": {},
		".leveldb": {}, ".wav": {}, ".rsh": {}, ".chromium": {}, ".rom": {}, ".ino": {}, ".pyi": {}, ".toml": {},
		".pyd": {}, ".pyx": {}, ".pdx": {}, ".code-snippets": {}, ".odt": {}, ".tex": {}, ".sty": {}, ".sasl": {},
		".crypt": {}, ".ed": {}, ".se": {}, ".prefs": {}, ".dynalink": {}, ".compiler": {}, ".attach": {}, ".sql": {},
		".rmi": {}, ".agent": {}, ".sqlite-shm": {}, ".sqlite-wal": {}, ".lproj": {}, ".scpt": {}, ".ans": {},
		".lua": {}, ".luac": {}, ".epub": {}, ".mobi": {}, ".numbers": {}, ".pages": {}, ".license": {}, ".lsf": {},
		".lsv": {}, ".gam": {}, ".lsb": {}, ".lsx": {}, ".ca-bundle": {}, ".adoc": {}, ".bc": {}, ".modulemap": {},
		".nashhorn": {}, ".shell": {}, ".management": {}, ".le": {}, ".ci": {}, ".ec": {}, ".snippets": {},
		".vim": {}, ".r": {}, ".tmpreferences": {}, ".nua": {}, ".icls": {}, ".fingerprint": {}, ".pb": {},
		".dmp": {}, ".tflite": {}, ".final": {}, ".mid": {}, ".pxd": {}, ".htc": {}, ".tmlanguage": {},
		".tmsnippet": {}, ".vmoptions": {}, ".plugins": {},

		// unsupported image types
		".tga": {}, ".ico": {}, ".pdf": {}, ".eps": {}, ".ai": {},
	}
	ext := filepath.Ext(strings.ToLower(filename))
	_, ok := imgExt[ext]
	if !ok {
		_, ok = skipExt[ext]
		if ok {
			return extSkipped
		}
		// anything with a number in it
		if must.OkOne(filepath.Match(".*[0-9]*", ext)) {
			return extSkipped
		}
		if must.OkOne(filepath.Match(".*ignore", ext)) {
			return extSkipped
		}
		if must.OkOne(filepath.Match(".*gz", ext)) {
			return extSkipped
		}
		if must.OkOne(filepath.Match(".*~*", ext)) {
			return extSkipped
		}
		if must.OkOne(filepath.Match(".xc*", ext)) {
			return extSkipped
		}
		return extUnknown
	}
	return extOK
}
