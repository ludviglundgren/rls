// Code generated by "stringer -type TagType -trimprefix TagType"; DO NOT EDIT.

package rls

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TagTypeWhitespace-0]
	_ = x[TagTypeDelim-1]
	_ = x[TagTypeText-2]
	_ = x[TagTypePlatform-3]
	_ = x[TagTypeArch-4]
	_ = x[TagTypeSource-5]
	_ = x[TagTypeResolution-6]
	_ = x[TagTypeCollection-7]
	_ = x[TagTypeDate-8]
	_ = x[TagTypeSeries-9]
	_ = x[TagTypeVersion-10]
	_ = x[TagTypeDisc-11]
	_ = x[TagTypeCodec-12]
	_ = x[TagTypeAudio-13]
	_ = x[TagTypeChannels-14]
	_ = x[TagTypeOther-15]
	_ = x[TagTypeCut-16]
	_ = x[TagTypeEdition-17]
	_ = x[TagTypeLanguage-18]
	_ = x[TagTypeSize-19]
	_ = x[TagTypeRegion-20]
	_ = x[TagTypeContainer-21]
	_ = x[TagTypeGenre-22]
	_ = x[TagTypeID-23]
	_ = x[TagTypeGroup-24]
	_ = x[TagTypeMeta-25]
	_ = x[TagTypeExt-26]
}

const _TagType_name = "WhitespaceDelimTextPlatformArchSourceResolutionCollectionDateSeriesVersionDiscCodecAudioChannelsOtherCutEditionLanguageSizeRegionContainerGenreIDGroupMetaExt"

var _TagType_index = [...]uint8{0, 10, 15, 19, 27, 31, 37, 47, 57, 61, 67, 74, 78, 83, 88, 96, 101, 104, 111, 119, 123, 129, 138, 143, 145, 150, 154, 157}

func (i TagType) String() string {
	if i < 0 || i >= TagType(len(_TagType_index)-1) {
		return "TagType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TagType_name[_TagType_index[i]:_TagType_index[i+1]]
}
