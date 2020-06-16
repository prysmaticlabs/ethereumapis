// Code generated by fastssz. DO NOT EDIT.
package eth

import (
	"fmt"

	ssz "github.com/ferranbt/fastssz"
)

var (
	errDivideInt           = fmt.Errorf("incorrect int divide")
	errListTooBig          = fmt.Errorf("incorrect list size, too big")
	errMarshalDynamicBytes = fmt.Errorf("incorrect dynamic bytes marshalling")
	errMarshalFixedBytes   = fmt.Errorf("incorrect fixed bytes marshalling")
	errMarshalList         = fmt.Errorf("incorrect vector list")
	errMarshalVector       = fmt.Errorf("incorrect vector marshalling")
	errOffset              = fmt.Errorf("incorrect offset")
	errSize                = fmt.Errorf("incorrect size")
)

// MarshalSSZ ssz marshals the BeaconBlock object
func (b *BeaconBlock) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, b.SizeSSZ())
	return b.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the BeaconBlock object to a target array
func (b *BeaconBlock) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error
	offset := int(84)

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, b.Slot)

	// Field (1) 'ProposerIndex'
	dst = ssz.MarshalUint64(dst, b.ProposerIndex)

	// Field (2) 'ParentRoot'
	if dst, err = ssz.MarshalFixedBytes(dst, b.ParentRoot, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (3) 'StateRoot'
	if dst, err = ssz.MarshalFixedBytes(dst, b.StateRoot, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Offset (4) 'Body'
	dst = ssz.WriteOffset(dst, offset)
	if b.Body == nil {
		b.Body = new(BeaconBlockBody)
	}
	offset += b.Body.SizeSSZ()

	// Field (4) 'Body'
	if dst, err = b.Body.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the BeaconBlock object
func (b *BeaconBlock) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 84 {
		return errSize
	}

	tail := buf
	var o4 uint64

	// Field (0) 'Slot'
	b.Slot = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ProposerIndex'
	b.ProposerIndex = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'ParentRoot'
	b.ParentRoot = append(b.ParentRoot, buf[16:48]...)

	// Field (3) 'StateRoot'
	b.StateRoot = append(b.StateRoot, buf[48:80]...)

	// Offset (4) 'Body'
	if o4 = ssz.ReadOffset(buf[80:84]); o4 > size {
		return errOffset
	}

	// Field (4) 'Body'
	{
		buf = tail[o4:]
		if b.Body == nil {
			b.Body = new(BeaconBlockBody)
		}
		if err = b.Body.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlock object
func (b *BeaconBlock) SizeSSZ() (size int) {
	size = 84

	// Field (4) 'Body'
	if b.Body == nil {
		b.Body = new(BeaconBlockBody)
	}
	size += b.Body.SizeSSZ()

	return
}

// MarshalSSZ ssz marshals the SignedBeaconBlock object
func (s *SignedBeaconBlock) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, s.SizeSSZ())
	return s.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the SignedBeaconBlock object to a target array
func (s *SignedBeaconBlock) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error
	offset := int(100)

	// Offset (0) 'Block'
	dst = ssz.WriteOffset(dst, offset)
	if s.Block == nil {
		s.Block = new(BeaconBlock)
	}
	offset += s.Block.SizeSSZ()

	// Field (1) 'Signature'
	if dst, err = ssz.MarshalFixedBytes(dst, s.Signature, 96); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (0) 'Block'
	if dst, err = s.Block.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the SignedBeaconBlock object
func (s *SignedBeaconBlock) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 100 {
		return errSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Block'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return errOffset
	}

	// Field (1) 'Signature'
	s.Signature = append(s.Signature, buf[4:100]...)

	// Field (0) 'Block'
	{
		buf = tail[o0:]
		if s.Block == nil {
			s.Block = new(BeaconBlock)
		}
		if err = s.Block.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedBeaconBlock object
func (s *SignedBeaconBlock) SizeSSZ() (size int) {
	size = 100

	// Field (0) 'Block'
	if s.Block == nil {
		s.Block = new(BeaconBlock)
	}
	size += s.Block.SizeSSZ()

	return
}

// MarshalSSZ ssz marshals the BeaconBlockBody object
func (b *BeaconBlockBody) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, b.SizeSSZ())
	return b.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the BeaconBlockBody object to a target array
func (b *BeaconBlockBody) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error
	offset := int(220)

	// Field (0) 'RandaoReveal'
	if dst, err = ssz.MarshalFixedBytes(dst, b.RandaoReveal, 96); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (1) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(Eth1Data)
	}
	if dst, err = b.Eth1Data.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	// Field (2) 'Graffiti'
	if dst, err = ssz.MarshalFixedBytes(dst, b.Graffiti, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Offset (3) 'ProposerSlashings'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.ProposerSlashings) * 416

	// Offset (4) 'AttesterSlashings'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.AttesterSlashings); ii++ {
		offset += 4
		offset += b.AttesterSlashings[ii].SizeSSZ()
	}

	// Offset (5) 'Attestations'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.Attestations); ii++ {
		offset += 4
		offset += b.Attestations[ii].SizeSSZ()
	}

	// Offset (6) 'Deposits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Deposits) * 1240

	// Offset (7) 'VoluntaryExits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.VoluntaryExits) * 112

	// Field (3) 'ProposerSlashings'
	if len(b.ProposerSlashings) > 16 {
		return nil, errMarshalList
	}
	for ii := 0; ii < len(b.ProposerSlashings); ii++ {
		if dst, err = b.ProposerSlashings[ii].MarshalSSZTo(dst); err != nil {
			return nil, err
		}
	}

	// Field (4) 'AttesterSlashings'
	if len(b.AttesterSlashings) > 2 {
		return nil, errMarshalList
	}
	{
		offset = 4 * len(b.AttesterSlashings)
		for ii := 0; ii < len(b.AttesterSlashings); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.AttesterSlashings[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.AttesterSlashings); ii++ {
		if dst, err = b.AttesterSlashings[ii].MarshalSSZTo(dst); err != nil {
			return nil, err
		}
	}

	// Field (5) 'Attestations'
	if len(b.Attestations) > 128 {
		return nil, errMarshalList
	}
	{
		offset = 4 * len(b.Attestations)
		for ii := 0; ii < len(b.Attestations); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.Attestations[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.Attestations); ii++ {
		if dst, err = b.Attestations[ii].MarshalSSZTo(dst); err != nil {
			return nil, err
		}
	}

	// Field (6) 'Deposits'
	if len(b.Deposits) > 16 {
		return nil, errMarshalList
	}
	for ii := 0; ii < len(b.Deposits); ii++ {
		if dst, err = b.Deposits[ii].MarshalSSZTo(dst); err != nil {
			return nil, err
		}
	}

	// Field (7) 'VoluntaryExits'
	if len(b.VoluntaryExits) > 16 {
		return nil, errMarshalList
	}
	for ii := 0; ii < len(b.VoluntaryExits); ii++ {
		if dst, err = b.VoluntaryExits[ii].MarshalSSZTo(dst); err != nil {
			return nil, err
		}
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockBody object
func (b *BeaconBlockBody) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 220 {
		return errSize
	}

	tail := buf
	var o3, o4, o5, o6, o7 uint64

	// Field (0) 'RandaoReveal'
	b.RandaoReveal = append(b.RandaoReveal, buf[0:96]...)

	// Field (1) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(Eth1Data)
	}
	if err = b.Eth1Data.UnmarshalSSZ(buf[96:168]); err != nil {
		return err
	}

	// Field (2) 'Graffiti'
	b.Graffiti = append(b.Graffiti, buf[168:200]...)

	// Offset (3) 'ProposerSlashings'
	if o3 = ssz.ReadOffset(buf[200:204]); o3 > size {
		return errOffset
	}

	// Offset (4) 'AttesterSlashings'
	if o4 = ssz.ReadOffset(buf[204:208]); o4 > size || o3 > o4 {
		return errOffset
	}

	// Offset (5) 'Attestations'
	if o5 = ssz.ReadOffset(buf[208:212]); o5 > size || o4 > o5 {
		return errOffset
	}

	// Offset (6) 'Deposits'
	if o6 = ssz.ReadOffset(buf[212:216]); o6 > size || o5 > o6 {
		return errOffset
	}

	// Offset (7) 'VoluntaryExits'
	if o7 = ssz.ReadOffset(buf[216:220]); o7 > size || o6 > o7 {
		return errOffset
	}

	// Field (3) 'ProposerSlashings'
	{
		buf = tail[o3:o4]
		num, ok := ssz.DivideInt(len(buf), 416)
		if !ok {
			return errDivideInt
		}
		if num > 16 {
			return errListTooBig
		}
		b.ProposerSlashings = make([]*ProposerSlashing, num)
		for ii := 0; ii < num; ii++ {
			if b.ProposerSlashings[ii] == nil {
				b.ProposerSlashings[ii] = new(ProposerSlashing)
			}
			if err = b.ProposerSlashings[ii].UnmarshalSSZ(buf[ii*416 : (ii+1)*416]); err != nil {
				return err
			}
		}
	}

	// Field (4) 'AttesterSlashings'
	{
		buf = tail[o4:o5]
		num, err := ssz.DecodeDynamicLength(buf, 2)
		if err != nil {
			return err
		}
		b.AttesterSlashings = make([]*AttesterSlashing, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.AttesterSlashings[indx] == nil {
				b.AttesterSlashings[indx] = new(AttesterSlashing)
			}
			if err = b.AttesterSlashings[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (5) 'Attestations'
	{
		buf = tail[o5:o6]
		num, err := ssz.DecodeDynamicLength(buf, 128)
		if err != nil {
			return err
		}
		b.Attestations = make([]*Attestation, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.Attestations[indx] == nil {
				b.Attestations[indx] = new(Attestation)
			}
			if err = b.Attestations[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (6) 'Deposits'
	{
		buf = tail[o6:o7]
		num, ok := ssz.DivideInt(len(buf), 1240)
		if !ok {
			return errDivideInt
		}
		if num > 16 {
			return errListTooBig
		}
		b.Deposits = make([]*Deposit, num)
		for ii := 0; ii < num; ii++ {
			if b.Deposits[ii] == nil {
				b.Deposits[ii] = new(Deposit)
			}
			if err = b.Deposits[ii].UnmarshalSSZ(buf[ii*1240 : (ii+1)*1240]); err != nil {
				return err
			}
		}
	}

	// Field (7) 'VoluntaryExits'
	{
		buf = tail[o7:]
		num, ok := ssz.DivideInt(len(buf), 112)
		if !ok {
			return errDivideInt
		}
		if num > 16 {
			return errListTooBig
		}
		b.VoluntaryExits = make([]*SignedVoluntaryExit, num)
		for ii := 0; ii < num; ii++ {
			if b.VoluntaryExits[ii] == nil {
				b.VoluntaryExits[ii] = new(SignedVoluntaryExit)
			}
			if err = b.VoluntaryExits[ii].UnmarshalSSZ(buf[ii*112 : (ii+1)*112]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockBody object
func (b *BeaconBlockBody) SizeSSZ() (size int) {
	size = 220

	// Field (3) 'ProposerSlashings'
	size += len(b.ProposerSlashings) * 416

	// Field (4) 'AttesterSlashings'
	for ii := 0; ii < len(b.AttesterSlashings); ii++ {
		size += 4
		size += b.AttesterSlashings[ii].SizeSSZ()
	}

	// Field (5) 'Attestations'
	for ii := 0; ii < len(b.Attestations); ii++ {
		size += 4
		size += b.Attestations[ii].SizeSSZ()
	}

	// Field (6) 'Deposits'
	size += len(b.Deposits) * 1240

	// Field (7) 'VoluntaryExits'
	size += len(b.VoluntaryExits) * 112

	return
}

// MarshalSSZ ssz marshals the ProposerSlashing object
func (p *ProposerSlashing) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, p.SizeSSZ())
	return p.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the ProposerSlashing object to a target array
func (p *ProposerSlashing) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error

	// Field (0) 'Header_1'
	if p.Header_1 == nil {
		p.Header_1 = new(SignedBeaconBlockHeader)
	}
	if dst, err = p.Header_1.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	// Field (1) 'Header_2'
	if p.Header_2 == nil {
		p.Header_2 = new(SignedBeaconBlockHeader)
	}
	if dst, err = p.Header_2.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the ProposerSlashing object
func (p *ProposerSlashing) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 416 {
		return errSize
	}

	// Field (0) 'Header_1'
	if p.Header_1 == nil {
		p.Header_1 = new(SignedBeaconBlockHeader)
	}
	if err = p.Header_1.UnmarshalSSZ(buf[0:208]); err != nil {
		return err
	}

	// Field (1) 'Header_2'
	if p.Header_2 == nil {
		p.Header_2 = new(SignedBeaconBlockHeader)
	}
	if err = p.Header_2.UnmarshalSSZ(buf[208:416]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ProposerSlashing object
func (p *ProposerSlashing) SizeSSZ() (size int) {
	size = 416
	return
}

// MarshalSSZ ssz marshals the AttesterSlashing object
func (a *AttesterSlashing) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, a.SizeSSZ())
	return a.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the AttesterSlashing object to a target array
func (a *AttesterSlashing) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error
	offset := int(8)

	// Offset (0) 'Attestation_1'
	dst = ssz.WriteOffset(dst, offset)
	if a.Attestation_1 == nil {
		a.Attestation_1 = new(IndexedAttestation)
	}
	offset += a.Attestation_1.SizeSSZ()

	// Offset (1) 'Attestation_2'
	dst = ssz.WriteOffset(dst, offset)
	if a.Attestation_2 == nil {
		a.Attestation_2 = new(IndexedAttestation)
	}
	offset += a.Attestation_2.SizeSSZ()

	// Field (0) 'Attestation_1'
	if dst, err = a.Attestation_1.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	// Field (1) 'Attestation_2'
	if dst, err = a.Attestation_2.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the AttesterSlashing object
func (a *AttesterSlashing) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return errSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'Attestation_1'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return errOffset
	}

	// Offset (1) 'Attestation_2'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return errOffset
	}

	// Field (0) 'Attestation_1'
	{
		buf = tail[o0:o1]
		if a.Attestation_1 == nil {
			a.Attestation_1 = new(IndexedAttestation)
		}
		if err = a.Attestation_1.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'Attestation_2'
	{
		buf = tail[o1:]
		if a.Attestation_2 == nil {
			a.Attestation_2 = new(IndexedAttestation)
		}
		if err = a.Attestation_2.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the AttesterSlashing object
func (a *AttesterSlashing) SizeSSZ() (size int) {
	size = 8

	// Field (0) 'Attestation_1'
	if a.Attestation_1 == nil {
		a.Attestation_1 = new(IndexedAttestation)
	}
	size += a.Attestation_1.SizeSSZ()

	// Field (1) 'Attestation_2'
	if a.Attestation_2 == nil {
		a.Attestation_2 = new(IndexedAttestation)
	}
	size += a.Attestation_2.SizeSSZ()

	return
}

// MarshalSSZ ssz marshals the Deposit object
func (d *Deposit) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, d.SizeSSZ())
	return d.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the Deposit object to a target array
func (d *Deposit) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error

	// Field (0) 'Proof'
	if len(d.Proof) != 33 {
		return nil, errMarshalVector
	}
	for ii := 0; ii < 33; ii++ {
		if dst, err = ssz.MarshalFixedBytes(dst, d.Proof[ii], 32); err != nil {
			return nil, errMarshalFixedBytes
		}
	}

	// Field (1) 'Data'
	if d.Data == nil {
		d.Data = new(Deposit_Data)
	}
	if dst, err = d.Data.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the Deposit object
func (d *Deposit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 1240 {
		return errSize
	}

	// Field (0) 'Proof'
	d.Proof = make([][]byte, 33)
	for ii := 0; ii < 33; ii++ {
		d.Proof[ii] = append(d.Proof[ii], buf[0:1056][ii*32:(ii+1)*32]...)
	}

	// Field (1) 'Data'
	if d.Data == nil {
		d.Data = new(Deposit_Data)
	}
	if err = d.Data.UnmarshalSSZ(buf[1056:1240]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Deposit object
func (d *Deposit) SizeSSZ() (size int) {
	size = 1240
	return
}

// MarshalSSZ ssz marshals the Deposit_Data object
func (d *Deposit_Data) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, d.SizeSSZ())
	return d.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the Deposit_Data object to a target array
func (d *Deposit_Data) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error

	// Field (0) 'PublicKey'
	if dst, err = ssz.MarshalFixedBytes(dst, d.PublicKey, 48); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (1) 'WithdrawalCredentials'
	if dst, err = ssz.MarshalFixedBytes(dst, d.WithdrawalCredentials, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (2) 'Amount'
	dst = ssz.MarshalUint64(dst, d.Amount)

	// Field (3) 'Signature'
	if dst, err = ssz.MarshalFixedBytes(dst, d.Signature, 96); err != nil {
		return nil, errMarshalFixedBytes
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the Deposit_Data object
func (d *Deposit_Data) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 184 {
		return errSize
	}

	// Field (0) 'PublicKey'
	d.PublicKey = append(d.PublicKey, buf[0:48]...)

	// Field (1) 'WithdrawalCredentials'
	d.WithdrawalCredentials = append(d.WithdrawalCredentials, buf[48:80]...)

	// Field (2) 'Amount'
	d.Amount = ssz.UnmarshallUint64(buf[80:88])

	// Field (3) 'Signature'
	d.Signature = append(d.Signature, buf[88:184]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Deposit_Data object
func (d *Deposit_Data) SizeSSZ() (size int) {
	size = 184
	return
}

// MarshalSSZ ssz marshals the VoluntaryExit object
func (v *VoluntaryExit) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, v.SizeSSZ())
	return v.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the VoluntaryExit object to a target array
func (v *VoluntaryExit) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error

	// Field (0) 'Epoch'
	dst = ssz.MarshalUint64(dst, v.Epoch)

	// Field (1) 'ValidatorIndex'
	dst = ssz.MarshalUint64(dst, v.ValidatorIndex)

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the VoluntaryExit object
func (v *VoluntaryExit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return errSize
	}

	// Field (0) 'Epoch'
	v.Epoch = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ValidatorIndex'
	v.ValidatorIndex = ssz.UnmarshallUint64(buf[8:16])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the VoluntaryExit object
func (v *VoluntaryExit) SizeSSZ() (size int) {
	size = 16
	return
}

// MarshalSSZ ssz marshals the SignedVoluntaryExit object
func (s *SignedVoluntaryExit) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, s.SizeSSZ())
	return s.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the SignedVoluntaryExit object to a target array
func (s *SignedVoluntaryExit) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error

	// Field (0) 'Exit'
	if s.Exit == nil {
		s.Exit = new(VoluntaryExit)
	}
	if dst, err = s.Exit.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	// Field (1) 'Signature'
	if dst, err = ssz.MarshalFixedBytes(dst, s.Signature, 96); err != nil {
		return nil, errMarshalFixedBytes
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the SignedVoluntaryExit object
func (s *SignedVoluntaryExit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 112 {
		return errSize
	}

	// Field (0) 'Exit'
	if s.Exit == nil {
		s.Exit = new(VoluntaryExit)
	}
	if err = s.Exit.UnmarshalSSZ(buf[0:16]); err != nil {
		return err
	}

	// Field (1) 'Signature'
	s.Signature = append(s.Signature, buf[16:112]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedVoluntaryExit object
func (s *SignedVoluntaryExit) SizeSSZ() (size int) {
	size = 112
	return
}

// MarshalSSZ ssz marshals the Eth1Data object
func (e *Eth1Data) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, e.SizeSSZ())
	return e.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the Eth1Data object to a target array
func (e *Eth1Data) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error

	// Field (0) 'DepositRoot'
	if dst, err = ssz.MarshalFixedBytes(dst, e.DepositRoot, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (1) 'DepositCount'
	dst = ssz.MarshalUint64(dst, e.DepositCount)

	// Field (2) 'BlockHash'
	if dst, err = ssz.MarshalFixedBytes(dst, e.BlockHash, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the Eth1Data object
func (e *Eth1Data) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 72 {
		return errSize
	}

	// Field (0) 'DepositRoot'
	e.DepositRoot = append(e.DepositRoot, buf[0:32]...)

	// Field (1) 'DepositCount'
	e.DepositCount = ssz.UnmarshallUint64(buf[32:40])

	// Field (2) 'BlockHash'
	e.BlockHash = append(e.BlockHash, buf[40:72]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Eth1Data object
func (e *Eth1Data) SizeSSZ() (size int) {
	size = 72
	return
}

// MarshalSSZ ssz marshals the BeaconBlockHeader object
func (b *BeaconBlockHeader) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, b.SizeSSZ())
	return b.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the BeaconBlockHeader object to a target array
func (b *BeaconBlockHeader) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, b.Slot)

	// Field (1) 'ProposerIndex'
	dst = ssz.MarshalUint64(dst, b.ProposerIndex)

	// Field (2) 'ParentRoot'
	if dst, err = ssz.MarshalFixedBytes(dst, b.ParentRoot, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (3) 'StateRoot'
	if dst, err = ssz.MarshalFixedBytes(dst, b.StateRoot, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (4) 'BodyRoot'
	if dst, err = ssz.MarshalFixedBytes(dst, b.BodyRoot, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockHeader object
func (b *BeaconBlockHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 112 {
		return errSize
	}

	// Field (0) 'Slot'
	b.Slot = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ProposerIndex'
	b.ProposerIndex = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'ParentRoot'
	b.ParentRoot = append(b.ParentRoot, buf[16:48]...)

	// Field (3) 'StateRoot'
	b.StateRoot = append(b.StateRoot, buf[48:80]...)

	// Field (4) 'BodyRoot'
	b.BodyRoot = append(b.BodyRoot, buf[80:112]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockHeader object
func (b *BeaconBlockHeader) SizeSSZ() (size int) {
	size = 112
	return
}

// MarshalSSZ ssz marshals the SignedBeaconBlockHeader object
func (s *SignedBeaconBlockHeader) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, s.SizeSSZ())
	return s.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the SignedBeaconBlockHeader object to a target array
func (s *SignedBeaconBlockHeader) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error

	// Field (0) 'Header'
	if s.Header == nil {
		s.Header = new(BeaconBlockHeader)
	}
	if dst, err = s.Header.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	// Field (1) 'Signature'
	if dst, err = ssz.MarshalFixedBytes(dst, s.Signature, 96); err != nil {
		return nil, errMarshalFixedBytes
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the SignedBeaconBlockHeader object
func (s *SignedBeaconBlockHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 208 {
		return errSize
	}

	// Field (0) 'Header'
	if s.Header == nil {
		s.Header = new(BeaconBlockHeader)
	}
	if err = s.Header.UnmarshalSSZ(buf[0:112]); err != nil {
		return err
	}

	// Field (1) 'Signature'
	s.Signature = append(s.Signature, buf[112:208]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedBeaconBlockHeader object
func (s *SignedBeaconBlockHeader) SizeSSZ() (size int) {
	size = 208
	return
}

// MarshalSSZ ssz marshals the IndexedAttestation object
func (i *IndexedAttestation) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, i.SizeSSZ())
	return i.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the IndexedAttestation object to a target array
func (i *IndexedAttestation) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error
	offset := int(228)

	// Offset (0) 'AttestingIndices'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(i.AttestingIndices) * 8

	// Field (1) 'Data'
	if i.Data == nil {
		i.Data = new(AttestationData)
	}
	if dst, err = i.Data.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	// Field (2) 'Signature'
	if dst, err = ssz.MarshalFixedBytes(dst, i.Signature, 96); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (0) 'AttestingIndices'
	if len(i.AttestingIndices) > 2048 {
		return nil, errMarshalList
	}
	for ii := 0; ii < len(i.AttestingIndices); ii++ {
		dst = ssz.MarshalUint64(dst, i.AttestingIndices[ii])
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the IndexedAttestation object
func (i *IndexedAttestation) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 228 {
		return errSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'AttestingIndices'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return errOffset
	}

	// Field (1) 'Data'
	if i.Data == nil {
		i.Data = new(AttestationData)
	}
	if err = i.Data.UnmarshalSSZ(buf[4:132]); err != nil {
		return err
	}

	// Field (2) 'Signature'
	i.Signature = append(i.Signature, buf[132:228]...)

	// Field (0) 'AttestingIndices'
	{
		buf = tail[o0:]
		num, ok := ssz.DivideInt(len(buf), 8)
		if !ok {
			return errDivideInt
		}
		if num > 2048 {
			return errListTooBig
		}
		i.AttestingIndices = ssz.ExtendUint64(i.AttestingIndices, num)
		for ii := 0; ii < num; ii++ {
			i.AttestingIndices[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the IndexedAttestation object
func (i *IndexedAttestation) SizeSSZ() (size int) {
	size = 228

	// Field (0) 'AttestingIndices'
	size += len(i.AttestingIndices) * 8

	return
}

// MarshalSSZ ssz marshals the Attestation object
func (a *Attestation) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, a.SizeSSZ())
	return a.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the Attestation object to a target array
func (a *Attestation) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error
	offset := int(228)

	// Offset (0) 'AggregationBits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(a.AggregationBits)

	// Field (1) 'Data'
	if a.Data == nil {
		a.Data = new(AttestationData)
	}
	if dst, err = a.Data.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	// Field (2) 'Signature'
	if dst, err = ssz.MarshalFixedBytes(dst, a.Signature, 96); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (0) 'AggregationBits'
	dst = append(dst, a.AggregationBits...)

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the Attestation object
func (a *Attestation) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 228 {
		return errSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'AggregationBits'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return errOffset
	}

	// Field (1) 'Data'
	if a.Data == nil {
		a.Data = new(AttestationData)
	}
	if err = a.Data.UnmarshalSSZ(buf[4:132]); err != nil {
		return err
	}

	// Field (2) 'Signature'
	a.Signature = append(a.Signature, buf[132:228]...)

	// Field (0) 'AggregationBits'
	{
		buf = tail[o0:]
		a.AggregationBits = append(a.AggregationBits, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Attestation object
func (a *Attestation) SizeSSZ() (size int) {
	size = 228

	// Field (0) 'AggregationBits'
	size += len(a.AggregationBits)

	return
}

// MarshalSSZ ssz marshals the AggregateAttestationAndProof object
func (a *AggregateAttestationAndProof) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, a.SizeSSZ())
	return a.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the AggregateAttestationAndProof object to a target array
func (a *AggregateAttestationAndProof) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error
	offset := int(108)

	// Field (0) 'AggregatorIndex'
	dst = ssz.MarshalUint64(dst, a.AggregatorIndex)

	// Offset (1) 'Aggregate'
	dst = ssz.WriteOffset(dst, offset)
	if a.Aggregate == nil {
		a.Aggregate = new(Attestation)
	}
	offset += a.Aggregate.SizeSSZ()

	// Field (2) 'SelectionProof'
	if dst, err = ssz.MarshalFixedBytes(dst, a.SelectionProof, 96); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (1) 'Aggregate'
	if dst, err = a.Aggregate.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the AggregateAttestationAndProof object
func (a *AggregateAttestationAndProof) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 108 {
		return errSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'AggregatorIndex'
	a.AggregatorIndex = ssz.UnmarshallUint64(buf[0:8])

	// Offset (1) 'Aggregate'
	if o1 = ssz.ReadOffset(buf[8:12]); o1 > size {
		return errOffset
	}

	// Field (2) 'SelectionProof'
	a.SelectionProof = append(a.SelectionProof, buf[12:108]...)

	// Field (1) 'Aggregate'
	{
		buf = tail[o1:]
		if a.Aggregate == nil {
			a.Aggregate = new(Attestation)
		}
		if err = a.Aggregate.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the AggregateAttestationAndProof object
func (a *AggregateAttestationAndProof) SizeSSZ() (size int) {
	size = 108

	// Field (1) 'Aggregate'
	if a.Aggregate == nil {
		a.Aggregate = new(Attestation)
	}
	size += a.Aggregate.SizeSSZ()

	return
}

// MarshalSSZ ssz marshals the SignedAggregateAttestationAndProof object
func (s *SignedAggregateAttestationAndProof) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, s.SizeSSZ())
	return s.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the SignedAggregateAttestationAndProof object to a target array
func (s *SignedAggregateAttestationAndProof) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error
	offset := int(100)

	// Offset (0) 'Message'
	dst = ssz.WriteOffset(dst, offset)
	if s.Message == nil {
		s.Message = new(AggregateAttestationAndProof)
	}
	offset += s.Message.SizeSSZ()

	// Field (1) 'Signature'
	if dst, err = ssz.MarshalFixedBytes(dst, s.Signature, 96); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (0) 'Message'
	if dst, err = s.Message.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the SignedAggregateAttestationAndProof object
func (s *SignedAggregateAttestationAndProof) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 100 {
		return errSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Message'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return errOffset
	}

	// Field (1) 'Signature'
	s.Signature = append(s.Signature, buf[4:100]...)

	// Field (0) 'Message'
	{
		buf = tail[o0:]
		if s.Message == nil {
			s.Message = new(AggregateAttestationAndProof)
		}
		if err = s.Message.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedAggregateAttestationAndProof object
func (s *SignedAggregateAttestationAndProof) SizeSSZ() (size int) {
	size = 100

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(AggregateAttestationAndProof)
	}
	size += s.Message.SizeSSZ()

	return
}

// MarshalSSZ ssz marshals the AttestationData object
func (a *AttestationData) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, a.SizeSSZ())
	return a.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the AttestationData object to a target array
func (a *AttestationData) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, a.Slot)

	// Field (1) 'CommitteeIndex'
	dst = ssz.MarshalUint64(dst, a.CommitteeIndex)

	// Field (2) 'BeaconBlockRoot'
	if dst, err = ssz.MarshalFixedBytes(dst, a.BeaconBlockRoot, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (3) 'Source'
	if a.Source == nil {
		a.Source = new(Checkpoint)
	}
	if dst, err = a.Source.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	// Field (4) 'Target'
	if a.Target == nil {
		a.Target = new(Checkpoint)
	}
	if dst, err = a.Target.MarshalSSZTo(dst); err != nil {
		return nil, err
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the AttestationData object
func (a *AttestationData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 128 {
		return errSize
	}

	// Field (0) 'Slot'
	a.Slot = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'CommitteeIndex'
	a.CommitteeIndex = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'BeaconBlockRoot'
	a.BeaconBlockRoot = append(a.BeaconBlockRoot, buf[16:48]...)

	// Field (3) 'Source'
	if a.Source == nil {
		a.Source = new(Checkpoint)
	}
	if err = a.Source.UnmarshalSSZ(buf[48:88]); err != nil {
		return err
	}

	// Field (4) 'Target'
	if a.Target == nil {
		a.Target = new(Checkpoint)
	}
	if err = a.Target.UnmarshalSSZ(buf[88:128]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the AttestationData object
func (a *AttestationData) SizeSSZ() (size int) {
	size = 128
	return
}

// MarshalSSZ ssz marshals the Checkpoint object
func (c *Checkpoint) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, c.SizeSSZ())
	return c.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the Checkpoint object to a target array
func (c *Checkpoint) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error

	// Field (0) 'Epoch'
	dst = ssz.MarshalUint64(dst, c.Epoch)

	// Field (1) 'Root'
	if dst, err = ssz.MarshalFixedBytes(dst, c.Root, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the Checkpoint object
func (c *Checkpoint) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 40 {
		return errSize
	}

	// Field (0) 'Epoch'
	c.Epoch = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Root'
	c.Root = append(c.Root, buf[8:40]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Checkpoint object
func (c *Checkpoint) SizeSSZ() (size int) {
	size = 40
	return
}

// MarshalSSZ ssz marshals the Validator object
func (v *Validator) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, v.SizeSSZ())
	return v.MarshalSSZTo(buf[:0])
}

// MarshalSSZTo ssz marshals the Validator object to a target array
func (v *Validator) MarshalSSZTo(dst []byte) ([]byte, error) {
	var err error

	// Field (0) 'PublicKey'
	if dst, err = ssz.MarshalFixedBytes(dst, v.PublicKey, 48); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (1) 'WithdrawalCredentials'
	if dst, err = ssz.MarshalFixedBytes(dst, v.WithdrawalCredentials, 32); err != nil {
		return nil, errMarshalFixedBytes
	}

	// Field (2) 'EffectiveBalance'
	dst = ssz.MarshalUint64(dst, v.EffectiveBalance)

	// Field (3) 'Slashed'
	dst = ssz.MarshalBool(dst, v.Slashed)

	// Field (4) 'ActivationEligibilityEpoch'
	dst = ssz.MarshalUint64(dst, v.ActivationEligibilityEpoch)

	// Field (5) 'ActivationEpoch'
	dst = ssz.MarshalUint64(dst, v.ActivationEpoch)

	// Field (6) 'ExitEpoch'
	dst = ssz.MarshalUint64(dst, v.ExitEpoch)

	// Field (7) 'WithdrawableEpoch'
	dst = ssz.MarshalUint64(dst, v.WithdrawableEpoch)

	return dst, err
}

// UnmarshalSSZ ssz unmarshals the Validator object
func (v *Validator) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 121 {
		return errSize
	}

	// Field (0) 'PublicKey'
	v.PublicKey = append(v.PublicKey, buf[0:48]...)

	// Field (1) 'WithdrawalCredentials'
	v.WithdrawalCredentials = append(v.WithdrawalCredentials, buf[48:80]...)

	// Field (2) 'EffectiveBalance'
	v.EffectiveBalance = ssz.UnmarshallUint64(buf[80:88])

	// Field (3) 'Slashed'
	v.Slashed = ssz.UnmarshalBool(buf[88:89])

	// Field (4) 'ActivationEligibilityEpoch'
	v.ActivationEligibilityEpoch = ssz.UnmarshallUint64(buf[89:97])

	// Field (5) 'ActivationEpoch'
	v.ActivationEpoch = ssz.UnmarshallUint64(buf[97:105])

	// Field (6) 'ExitEpoch'
	v.ExitEpoch = ssz.UnmarshallUint64(buf[105:113])

	// Field (7) 'WithdrawableEpoch'
	v.WithdrawableEpoch = ssz.UnmarshallUint64(buf[113:121])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Validator object
func (v *Validator) SizeSSZ() (size int) {
	size = 121
	return
}
