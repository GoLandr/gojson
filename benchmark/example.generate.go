// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY go-fish/gojson.
// ************************************************************

package benchmark

import (
	backend "github.com/go-fish/gojson/backend"
	errors "github.com/go-fish/gojson/errors"
)

func (c *CBAvatar) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()

	enc.WriteByte('{')
	enc.EncodeKeyString("url", c.Url)

	enc.WriteByte('}')
	data := enc.Bytes()
	enc.Release()

	return data, nil
}

func (c *CBAvatar) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder()
	dec.SetUnsafeData(data)

	if char := dec.NextChar(); char == 'n' {
		if err := dec.AssetNull(); err != nil {
			return err
		}

	} else if char != '{' {
		return errors.NewParseError(dec.Char(), dec.Cursor())
	} else {
		dec.Next()
		if dec.IsObjectClose() {
			return nil
		} else {
			for objdc8155ea73453729 := 1; objdc8155ea73453729 > 0; {
				keye447d0b30f72b9a6, err := dec.NextKey()
				if err != nil {
					return err
				}

				switch keye447d0b30f72b9a6 {
				case "url":
					valued7e76af3818193f6, err := dec.DecodeString()
					if err != nil {
						return err
					}

					c.Url = valued7e76af3818193f6

				default:
					if err := dec.SkipValue(); err != nil {
						return err
					}
				}
				if dec.IsObjectClose() {
					objdc8155ea73453729--
				}
			}
		}
	}
	dec.Release()

	return nil
}

func (c *CBGithub) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()

	enc.WriteByte('{')
	enc.EncodeKeyInt("followers", c.Followers)

	enc.WriteByte('}')
	data := enc.Bytes()
	enc.Release()

	return data, nil
}

func (c *CBGithub) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder()
	dec.SetUnsafeData(data)

	if char := dec.NextChar(); char == 'n' {
		if err := dec.AssetNull(); err != nil {
			return err
		}

	} else if char != '{' {
		return errors.NewParseError(dec.Char(), dec.Cursor())
	} else {
		dec.Next()
		if dec.IsObjectClose() {
			return nil
		} else {
			for obj31233a73a5f6a14f := 1; obj31233a73a5f6a14f > 0; {
				key605e9c2f08de0add, err := dec.NextKey()
				if err != nil {
					return err
				}

				switch key605e9c2f08de0add {
				case "followers":
					valuecb68fac2bc2789c4, err := dec.DecodeInt()
					if err != nil {
						return err
					}

					c.Followers = valuecb68fac2bc2789c4

				default:
					if err := dec.SkipValue(); err != nil {
						return err
					}
				}
				if dec.IsObjectClose() {
					obj31233a73a5f6a14f--
				}
			}
		}
	}
	dec.Release()

	return nil
}

func (c *CBGravatar) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()

	enc.WriteByte('{')
	enc.WriteKey("avatars")
	if len(c.Avatars) == 0 {
		enc.WriteNull()
	} else {
		enc.WriteByte('[')
		for _, valueb2f627e2e376cee2 := range c.Avatars {
			enc.WriteComma()
			if valueb2f627e2e376cee2 == nil {
				enc.WriteNull()
			} else {
				enc.WriteByte('{')
				enc.EncodeKeyString("url", valueb2f627e2e376cee2.Url)

				enc.WriteByte('}')
			}
		}

		enc.WriteByte(']')
	}
	enc.WriteByte('}')
	data := enc.Bytes()
	enc.Release()

	return data, nil
}

func (c *CBGravatar) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder()
	dec.SetUnsafeData(data)

	if char := dec.NextChar(); char == 'n' {
		if err := dec.AssetNull(); err != nil {
			return err
		}

	} else if char != '{' {
		return errors.NewParseError(dec.Char(), dec.Cursor())
	} else {
		dec.Next()
		if dec.IsObjectClose() {
			return nil
		} else {
			for objf7aefa4f83a4a00a := 1; objf7aefa4f83a4a00a > 0; {
				keyb8651e19abf08f60, err := dec.NextKey()
				if err != nil {
					return err
				}

				switch keyb8651e19abf08f60 {
				case "avatars":
					if char := dec.NextChar(); char == 'n' {
						if err := dec.AssetNull(); err != nil {
							return err
						}

						c.Avatars = nil
					} else if char != '[' {
						return errors.NewParseError(dec.Char(), dec.Cursor())
					} else {
						dec.Next()
						if dec.IsArrayClose() {
							c.Avatars = nil
						} else {
							if c.Avatars == nil {
								c.Avatars = make(Avatars, 0, 8)
							}

							for arrayb9acef45a7d35608 := 1; arrayb9acef45a7d35608 > 0; {
								var value2d5e32356098d4d4 *CBAvatar
								if dec.IsNull() {
									value2d5e32356098d4d4 = nil
								} else {
									value2d5e32356098d4d4 = new(CBAvatar)

									if char := dec.NextChar(); char == 'n' {
										if err := dec.AssetNull(); err != nil {
											return err
										}

									} else if char != '{' {
										return errors.NewParseError(dec.Char(), dec.Cursor())
									} else {
										dec.Next()
										if dec.IsObjectClose() {
											return nil
										} else {
											for obj9b96a580b6418388 := 1; obj9b96a580b6418388 > 0; {
												key2d6d317c5d5ed1f8, err := dec.NextKey()
												if err != nil {
													return err
												}

												switch key2d6d317c5d5ed1f8 {
												case "url":
													valuea8d008128f2e2483, err := dec.DecodeString()
													if err != nil {
														return err
													}

													value2d5e32356098d4d4.Url = valuea8d008128f2e2483

												default:
													if err := dec.SkipValue(); err != nil {
														return err
													}
												}
												if dec.IsObjectClose() {
													obj9b96a580b6418388--
												}
											}
										}
									}
								}
								if value2d5e32356098d4d4 != nil {
									c.Avatars = append(c.Avatars, value2d5e32356098d4d4)
								}
								if dec.IsArrayClose() {
									arrayb9acef45a7d35608--
								}
							}
						}
					}

				default:
					if err := dec.SkipValue(); err != nil {
						return err
					}
				}
				if dec.IsObjectClose() {
					objf7aefa4f83a4a00a--
				}
			}
		}
	}
	dec.Release()

	return nil
}

func (c *CBName) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()

	enc.WriteByte('{')
	enc.EncodeKeyString("fullName", c.FullName)

	enc.WriteByte('}')
	data := enc.Bytes()
	enc.Release()

	return data, nil
}

func (c *CBName) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder()
	dec.SetUnsafeData(data)

	if char := dec.NextChar(); char == 'n' {
		if err := dec.AssetNull(); err != nil {
			return err
		}

	} else if char != '{' {
		return errors.NewParseError(dec.Char(), dec.Cursor())
	} else {
		dec.Next()
		if dec.IsObjectClose() {
			return nil
		} else {
			for objf277d3571d47abf0 := 1; objf277d3571d47abf0 > 0; {
				key5c2452b721d1d482, err := dec.NextKey()
				if err != nil {
					return err
				}

				switch key5c2452b721d1d482 {
				case "fullName":
					valueec56040faa3f4b54, err := dec.DecodeString()
					if err != nil {
						return err
					}

					c.FullName = valueec56040faa3f4b54

				default:
					if err := dec.SkipValue(); err != nil {
						return err
					}
				}
				if dec.IsObjectClose() {
					objf277d3571d47abf0--
				}
			}
		}
	}
	dec.Release()

	return nil
}

func (c *CBPerson) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()

	enc.WriteByte('{')
	if c.Name != nil {
		enc.WriteKey("name")
		enc.WriteByte('{')
		enc.EncodeKeyString("fullName", c.Name.FullName)

		enc.WriteByte('}')
	}
	if c.Github != nil {
		enc.WriteKey("github")
		enc.WriteByte('{')
		enc.EncodeKeyInt("followers", c.Github.Followers)

		enc.WriteByte('}')
	}
	if c.Gravatar != nil {
		enc.WriteKey("gravatar")
		enc.WriteByte('{')
		enc.WriteKey("avatars")
		if len(c.Gravatar.Avatars) == 0 {
			enc.WriteNull()
		} else {
			enc.WriteByte('[')
			for _, value85c75f312cdb570c := range c.Gravatar.Avatars {
				enc.WriteComma()
				if value85c75f312cdb570c == nil {
					enc.WriteNull()
				} else {
					enc.WriteByte('{')
					enc.EncodeKeyString("url", value85c75f312cdb570c.Url)

					enc.WriteByte('}')
				}
			}

			enc.WriteByte(']')
		}
		enc.WriteByte('}')
	}
	enc.WriteByte('}')
	data := enc.Bytes()
	enc.Release()

	return data, nil
}

func (c *CBPerson) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder()
	dec.SetUnsafeData(data)

	if char := dec.NextChar(); char == 'n' {
		if err := dec.AssetNull(); err != nil {
			return err
		}

	} else if char != '{' {
		return errors.NewParseError(dec.Char(), dec.Cursor())
	} else {
		dec.Next()
		if dec.IsObjectClose() {
			return nil
		} else {
			for obj27f50ee5dc798e77 := 1; obj27f50ee5dc798e77 > 0; {
				key07e98cff2fbda131, err := dec.NextKey()
				if err != nil {
					return err
				}

				switch key07e98cff2fbda131 {
				case "name":
					if dec.IsNull() {
						c.Name = nil
					} else if !dec.IsObjectOpen() {
						return errors.NewParseError(dec.Char(), dec.Cursor())
					} else {
						if dec.IsObjectClose() {
							c.Name = nil
						} else {
							if c.Name == nil {
								c.Name = new(CBName)
							}

							for obj3c810279cec13618 := 1; obj3c810279cec13618 > 0; {
								key69452d12e09093c9, err := dec.NextKey()
								if err != nil {
									return err
								}

								switch key69452d12e09093c9 {
								case "fullName":
									value5baeeb9d79eea705, err := dec.DecodeString()
									if err != nil {
										return err
									}

									c.Name.FullName = value5baeeb9d79eea705

								default:
									if err := dec.SkipValue(); err != nil {
										return err
									}
								}
								if dec.IsObjectClose() {
									obj3c810279cec13618--
								}
							}
						}
					}

				case "github":
					if dec.IsNull() {
						c.Github = nil
					} else if !dec.IsObjectOpen() {
						return errors.NewParseError(dec.Char(), dec.Cursor())
					} else {
						if dec.IsObjectClose() {
							c.Github = nil
						} else {
							if c.Github == nil {
								c.Github = new(CBGithub)
							}

							for obj9ec925be3904b8d8 := 1; obj9ec925be3904b8d8 > 0; {
								keyd09334ae6a5d7ca2, err := dec.NextKey()
								if err != nil {
									return err
								}

								switch keyd09334ae6a5d7ca2 {
								case "followers":
									value68dcc2bb1ef27925, err := dec.DecodeInt()
									if err != nil {
										return err
									}

									c.Github.Followers = value68dcc2bb1ef27925

								default:
									if err := dec.SkipValue(); err != nil {
										return err
									}
								}
								if dec.IsObjectClose() {
									obj9ec925be3904b8d8--
								}
							}
						}
					}

				case "gravatar":
					if dec.IsNull() {
						c.Gravatar = nil
					} else if !dec.IsObjectOpen() {
						return errors.NewParseError(dec.Char(), dec.Cursor())
					} else {
						if dec.IsObjectClose() {
							c.Gravatar = nil
						} else {
							if c.Gravatar == nil {
								c.Gravatar = new(CBGravatar)
							}

							for objbb21b0a87fa969b8 := 1; objbb21b0a87fa969b8 > 0; {
								keybde68e76371b0edc, err := dec.NextKey()
								if err != nil {
									return err
								}

								switch keybde68e76371b0edc {
								case "avatars":
									if char := dec.NextChar(); char == 'n' {
										if err := dec.AssetNull(); err != nil {
											return err
										}

										c.Gravatar.Avatars = nil
									} else if char != '[' {
										return errors.NewParseError(dec.Char(), dec.Cursor())
									} else {
										dec.Next()
										if dec.IsArrayClose() {
											c.Gravatar.Avatars = nil
										} else {
											if c.Gravatar.Avatars == nil {
												c.Gravatar.Avatars = make(Avatars, 0, 8)
											}

											for arrayc8772200c60d8988 := 1; arrayc8772200c60d8988 > 0; {
												var valued03ce19709e1368d *CBAvatar
												if dec.IsNull() {
													valued03ce19709e1368d = nil
												} else {
													valued03ce19709e1368d = new(CBAvatar)

													if char := dec.NextChar(); char == 'n' {
														if err := dec.AssetNull(); err != nil {
															return err
														}

													} else if char != '{' {
														return errors.NewParseError(dec.Char(), dec.Cursor())
													} else {
														dec.Next()
														if dec.IsObjectClose() {
															return nil
														} else {
															for obj2949f2f705e3c87a := 1; obj2949f2f705e3c87a > 0; {
																keyb07c27752ebc8e07, err := dec.NextKey()
																if err != nil {
																	return err
																}

																switch keyb07c27752ebc8e07 {
																case "url":
																	value09c8a8f9e32b68cf, err := dec.DecodeString()
																	if err != nil {
																		return err
																	}

																	valued03ce19709e1368d.Url = value09c8a8f9e32b68cf

																default:
																	if err := dec.SkipValue(); err != nil {
																		return err
																	}
																}
																if dec.IsObjectClose() {
																	obj2949f2f705e3c87a--
																}
															}
														}
													}
												}
												if valued03ce19709e1368d != nil {
													c.Gravatar.Avatars = append(c.Gravatar.Avatars, valued03ce19709e1368d)
												}
												if dec.IsArrayClose() {
													arrayc8772200c60d8988--
												}
											}
										}
									}

								default:
									if err := dec.SkipValue(); err != nil {
										return err
									}
								}
								if dec.IsObjectClose() {
									objbb21b0a87fa969b8--
								}
							}
						}
					}

				default:
					if err := dec.SkipValue(); err != nil {
						return err
					}
				}
				if dec.IsObjectClose() {
					obj27f50ee5dc798e77--
				}
			}
		}
	}
	dec.Release()

	return nil
}

func (d *DSTopic) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()

	enc.WriteByte('{')
	enc.EncodeKeyInt("id", d.Id)

	enc.EncodeKeyString("slug", d.Slug)

	enc.WriteByte('}')
	data := enc.Bytes()
	enc.Release()

	return data, nil
}

func (d *DSTopic) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder()
	dec.SetUnsafeData(data)

	if char := dec.NextChar(); char == 'n' {
		if err := dec.AssetNull(); err != nil {
			return err
		}

	} else if char != '{' {
		return errors.NewParseError(dec.Char(), dec.Cursor())
	} else {
		dec.Next()
		if dec.IsObjectClose() {
			return nil
		} else {
			for objaf47527b2844386d := 1; objaf47527b2844386d > 0; {
				key5a709f17747e11ab, err := dec.NextKey()
				if err != nil {
					return err
				}

				switch key5a709f17747e11ab {
				case "id":
					valuee8bfe225689f3d38, err := dec.DecodeInt()
					if err != nil {
						return err
					}

					d.Id = valuee8bfe225689f3d38

				case "slug":
					value9ccde299b4a5cee9, err := dec.DecodeString()
					if err != nil {
						return err
					}

					d.Slug = value9ccde299b4a5cee9

				default:
					if err := dec.SkipValue(); err != nil {
						return err
					}
				}
				if dec.IsObjectClose() {
					objaf47527b2844386d--
				}
			}
		}
	}
	dec.Release()

	return nil
}

func (d *DSTopicsList) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()

	enc.WriteByte('{')
	enc.WriteKey("topics")
	if len(d.Topics) == 0 {
		enc.WriteNull()
	} else {
		enc.WriteByte('[')
		for _, valuee7bfa9be0fd4ce1b := range d.Topics {
			enc.WriteComma()
			if valuee7bfa9be0fd4ce1b == nil {
				enc.WriteNull()
			} else {
				enc.WriteByte('{')
				enc.EncodeKeyInt("id", valuee7bfa9be0fd4ce1b.Id)

				enc.EncodeKeyString("slug", valuee7bfa9be0fd4ce1b.Slug)

				enc.WriteByte('}')
			}
		}

		enc.WriteByte(']')
	}
	enc.EncodeKeyString("more_topics_url", d.MoreTopicsUrl)

	enc.WriteByte('}')
	data := enc.Bytes()
	enc.Release()

	return data, nil
}

func (d *DSTopicsList) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder()
	dec.SetUnsafeData(data)

	if char := dec.NextChar(); char == 'n' {
		if err := dec.AssetNull(); err != nil {
			return err
		}

	} else if char != '{' {
		return errors.NewParseError(dec.Char(), dec.Cursor())
	} else {
		dec.Next()
		if dec.IsObjectClose() {
			return nil
		} else {
			for obj6e7a5039fab17b48 := 1; obj6e7a5039fab17b48 > 0; {
				key9211e413371d82b8, err := dec.NextKey()
				if err != nil {
					return err
				}

				switch key9211e413371d82b8 {
				case "topics":
					if char := dec.NextChar(); char == 'n' {
						if err := dec.AssetNull(); err != nil {
							return err
						}

						d.Topics = nil
					} else if char != '[' {
						return errors.NewParseError(dec.Char(), dec.Cursor())
					} else {
						dec.Next()
						if dec.IsArrayClose() {
							d.Topics = nil
						} else {
							if d.Topics == nil {
								d.Topics = make(DSTopics, 0, 8)
							}

							for arraya0ab29cd417e1f26 := 1; arraya0ab29cd417e1f26 > 0; {
								var valueeef8244158da9fe5 *DSTopic
								if dec.IsNull() {
									valueeef8244158da9fe5 = nil
								} else {
									valueeef8244158da9fe5 = new(DSTopic)

									if char := dec.NextChar(); char == 'n' {
										if err := dec.AssetNull(); err != nil {
											return err
										}

									} else if char != '{' {
										return errors.NewParseError(dec.Char(), dec.Cursor())
									} else {
										dec.Next()
										if dec.IsObjectClose() {
											return nil
										} else {
											for objc433ff5ae434cbe8 := 1; objc433ff5ae434cbe8 > 0; {
												key526b5f81e2ae9efd, err := dec.NextKey()
												if err != nil {
													return err
												}

												switch key526b5f81e2ae9efd {
												case "id":
													valuef17f36a7593e5774, err := dec.DecodeInt()
													if err != nil {
														return err
													}

													valueeef8244158da9fe5.Id = valuef17f36a7593e5774

												case "slug":
													value80e985f75d963547, err := dec.DecodeString()
													if err != nil {
														return err
													}

													valueeef8244158da9fe5.Slug = value80e985f75d963547

												default:
													if err := dec.SkipValue(); err != nil {
														return err
													}
												}
												if dec.IsObjectClose() {
													objc433ff5ae434cbe8--
												}
											}
										}
									}
								}
								if valueeef8244158da9fe5 != nil {
									d.Topics = append(d.Topics, valueeef8244158da9fe5)
								}
								if dec.IsArrayClose() {
									arraya0ab29cd417e1f26--
								}
							}
						}
					}

				case "more_topics_url":
					valued0e4d838d0af60bd, err := dec.DecodeString()
					if err != nil {
						return err
					}

					d.MoreTopicsUrl = valued0e4d838d0af60bd

				default:
					if err := dec.SkipValue(); err != nil {
						return err
					}
				}
				if dec.IsObjectClose() {
					obj6e7a5039fab17b48--
				}
			}
		}
	}
	dec.Release()

	return nil
}

func (d *DSUser) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()

	enc.WriteByte('{')
	enc.EncodeKeyString("username", d.Username)

	enc.WriteByte('}')
	data := enc.Bytes()
	enc.Release()

	return data, nil
}

func (d *DSUser) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder()
	dec.SetUnsafeData(data)

	if char := dec.NextChar(); char == 'n' {
		if err := dec.AssetNull(); err != nil {
			return err
		}

	} else if char != '{' {
		return errors.NewParseError(dec.Char(), dec.Cursor())
	} else {
		dec.Next()
		if dec.IsObjectClose() {
			return nil
		} else {
			for obj9749416861de4852 := 1; obj9749416861de4852 > 0; {
				key2daefd87e960076f, err := dec.NextKey()
				if err != nil {
					return err
				}

				switch key2daefd87e960076f {
				case "username":
					value4fd38e192c55d6c7, err := dec.DecodeString()
					if err != nil {
						return err
					}

					d.Username = value4fd38e192c55d6c7

				default:
					if err := dec.SkipValue(); err != nil {
						return err
					}
				}
				if dec.IsObjectClose() {
					obj9749416861de4852--
				}
			}
		}
	}
	dec.Release()

	return nil
}

func (l *LargePayload) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()

	enc.WriteByte('{')
	enc.WriteKey("users")
	if len(l.Users) == 0 {
		enc.WriteNull()
	} else {
		enc.WriteByte('[')
		for _, value379e8a31475696ab := range l.Users {
			enc.WriteComma()
			if value379e8a31475696ab == nil {
				enc.WriteNull()
			} else {
				enc.WriteByte('{')
				enc.EncodeKeyString("username", value379e8a31475696ab.Username)

				enc.WriteByte('}')
			}
		}

		enc.WriteByte(']')
	}
	if l.Topics != nil {
		enc.WriteKey("topics")
		enc.WriteByte('{')
		enc.WriteKey("topics")
		if len(l.Topics.Topics) == 0 {
			enc.WriteNull()
		} else {
			enc.WriteByte('[')
			for _, value17f978a64313f083 := range l.Topics.Topics {
				enc.WriteComma()
				if value17f978a64313f083 == nil {
					enc.WriteNull()
				} else {
					enc.WriteByte('{')
					enc.EncodeKeyInt("id", value17f978a64313f083.Id)

					enc.EncodeKeyString("slug", value17f978a64313f083.Slug)

					enc.WriteByte('}')
				}
			}

			enc.WriteByte(']')
		}
		enc.EncodeKeyString("more_topics_url", l.Topics.MoreTopicsUrl)

		enc.WriteByte('}')
	}
	enc.WriteByte('}')
	data := enc.Bytes()
	enc.Release()

	return data, nil
}

func (l *LargePayload) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder()
	dec.SetUnsafeData(data)

	if char := dec.NextChar(); char == 'n' {
		if err := dec.AssetNull(); err != nil {
			return err
		}

	} else if char != '{' {
		return errors.NewParseError(dec.Char(), dec.Cursor())
	} else {
		dec.Next()
		if dec.IsObjectClose() {
			return nil
		} else {
			for obj156f5ef00a3ed054 := 1; obj156f5ef00a3ed054 > 0; {
				key72a980d57b6e39fa, err := dec.NextKey()
				if err != nil {
					return err
				}

				switch key72a980d57b6e39fa {
				case "users":
					if char := dec.NextChar(); char == 'n' {
						if err := dec.AssetNull(); err != nil {
							return err
						}

						l.Users = nil
					} else if char != '[' {
						return errors.NewParseError(dec.Char(), dec.Cursor())
					} else {
						dec.Next()
						if dec.IsArrayClose() {
							l.Users = nil
						} else {
							if l.Users == nil {
								l.Users = make(DSUsers, 0, 8)
							}

							for array05c44c2a6b0c1801 := 1; array05c44c2a6b0c1801 > 0; {
								var value1193038c94315cef *DSUser
								if dec.IsNull() {
									value1193038c94315cef = nil
								} else {
									value1193038c94315cef = new(DSUser)

									if char := dec.NextChar(); char == 'n' {
										if err := dec.AssetNull(); err != nil {
											return err
										}

									} else if char != '{' {
										return errors.NewParseError(dec.Char(), dec.Cursor())
									} else {
										dec.Next()
										if dec.IsObjectClose() {
											return nil
										} else {
											for obj39e38fe8376a5872 := 1; obj39e38fe8376a5872 > 0; {
												key7eb7d84022692e1e, err := dec.NextKey()
												if err != nil {
													return err
												}

												switch key7eb7d84022692e1e {
												case "username":
													value64e4ee14d2b3e23e, err := dec.DecodeString()
													if err != nil {
														return err
													}

													value1193038c94315cef.Username = value64e4ee14d2b3e23e

												default:
													if err := dec.SkipValue(); err != nil {
														return err
													}
												}
												if dec.IsObjectClose() {
													obj39e38fe8376a5872--
												}
											}
										}
									}
								}
								if value1193038c94315cef != nil {
									l.Users = append(l.Users, value1193038c94315cef)
								}
								if dec.IsArrayClose() {
									array05c44c2a6b0c1801--
								}
							}
						}
					}

				case "topics":
					if dec.IsNull() {
						l.Topics = nil
					} else if !dec.IsObjectOpen() {
						return errors.NewParseError(dec.Char(), dec.Cursor())
					} else {
						if dec.IsObjectClose() {
							l.Topics = nil
						} else {
							if l.Topics == nil {
								l.Topics = new(DSTopicsList)
							}

							for obj92c1c776f4a77147 := 1; obj92c1c776f4a77147 > 0; {
								key2698d3b085335af4, err := dec.NextKey()
								if err != nil {
									return err
								}

								switch key2698d3b085335af4 {
								case "topics":
									if char := dec.NextChar(); char == 'n' {
										if err := dec.AssetNull(); err != nil {
											return err
										}

										l.Topics.Topics = nil
									} else if char != '[' {
										return errors.NewParseError(dec.Char(), dec.Cursor())
									} else {
										dec.Next()
										if dec.IsArrayClose() {
											l.Topics.Topics = nil
										} else {
											if l.Topics.Topics == nil {
												l.Topics.Topics = make(DSTopics, 0, 8)
											}

											for arrayb42759d5ae6aa24e := 1; arrayb42759d5ae6aa24e > 0; {
												var value8a99ae44368ebe88 *DSTopic
												if dec.IsNull() {
													value8a99ae44368ebe88 = nil
												} else {
													value8a99ae44368ebe88 = new(DSTopic)

													if char := dec.NextChar(); char == 'n' {
														if err := dec.AssetNull(); err != nil {
															return err
														}

													} else if char != '{' {
														return errors.NewParseError(dec.Char(), dec.Cursor())
													} else {
														dec.Next()
														if dec.IsObjectClose() {
															return nil
														} else {
															for obj9dea6274c0003751 := 1; obj9dea6274c0003751 > 0; {
																key0884332ded0a59a9, err := dec.NextKey()
																if err != nil {
																	return err
																}

																switch key0884332ded0a59a9 {
																case "id":
																	value6c5fcdcd3ffdcc28, err := dec.DecodeInt()
																	if err != nil {
																		return err
																	}

																	value8a99ae44368ebe88.Id = value6c5fcdcd3ffdcc28

																case "slug":
																	value2cebc2e0cd7b7c68, err := dec.DecodeString()
																	if err != nil {
																		return err
																	}

																	value8a99ae44368ebe88.Slug = value2cebc2e0cd7b7c68

																default:
																	if err := dec.SkipValue(); err != nil {
																		return err
																	}
																}
																if dec.IsObjectClose() {
																	obj9dea6274c0003751--
																}
															}
														}
													}
												}
												if value8a99ae44368ebe88 != nil {
													l.Topics.Topics = append(l.Topics.Topics, value8a99ae44368ebe88)
												}
												if dec.IsArrayClose() {
													arrayb42759d5ae6aa24e--
												}
											}
										}
									}

								case "more_topics_url":
									valuef9a892ec5e536424, err := dec.DecodeString()
									if err != nil {
										return err
									}

									l.Topics.MoreTopicsUrl = valuef9a892ec5e536424

								default:
									if err := dec.SkipValue(); err != nil {
										return err
									}
								}
								if dec.IsObjectClose() {
									obj92c1c776f4a77147--
								}
							}
						}
					}

				default:
					if err := dec.SkipValue(); err != nil {
						return err
					}
				}
				if dec.IsObjectClose() {
					obj156f5ef00a3ed054--
				}
			}
		}
	}
	dec.Release()

	return nil
}

func (m *MediumPayload) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()

	enc.WriteByte('{')
	if m.Person != nil {
		enc.WriteKey("person")
		enc.WriteByte('{')
		if m.Person.Name != nil {
			enc.WriteKey("name")
			enc.WriteByte('{')
			enc.EncodeKeyString("fullName", m.Person.Name.FullName)

			enc.WriteByte('}')
		}
		if m.Person.Github != nil {
			enc.WriteKey("github")
			enc.WriteByte('{')
			enc.EncodeKeyInt("followers", m.Person.Github.Followers)

			enc.WriteByte('}')
		}
		if m.Person.Gravatar != nil {
			enc.WriteKey("gravatar")
			enc.WriteByte('{')
			enc.WriteKey("avatars")
			if len(m.Person.Gravatar.Avatars) == 0 {
				enc.WriteNull()
			} else {
				enc.WriteByte('[')
				for _, value5049c95a2d9df4db := range m.Person.Gravatar.Avatars {
					enc.WriteComma()
					if value5049c95a2d9df4db == nil {
						enc.WriteNull()
					} else {
						enc.WriteByte('{')
						enc.EncodeKeyString("url", value5049c95a2d9df4db.Url)

						enc.WriteByte('}')
					}
				}

				enc.WriteByte(']')
			}
			enc.WriteByte('}')
		}
		enc.WriteByte('}')
	}
	if m.Company != "" {
		enc.EncodeKeyString("company", m.Company)
	}

	enc.WriteByte('}')
	data := enc.Bytes()
	enc.Release()

	return data, nil
}

func (m *MediumPayload) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder()
	dec.SetUnsafeData(data)

	if char := dec.NextChar(); char == 'n' {
		if err := dec.AssetNull(); err != nil {
			return err
		}

	} else if char != '{' {
		return errors.NewParseError(dec.Char(), dec.Cursor())
	} else {
		dec.Next()
		if dec.IsObjectClose() {
			return nil
		} else {
			for obja785f5615ce785ac := 1; obja785f5615ce785ac > 0; {
				keyd108ca1e39974529, err := dec.NextKey()
				if err != nil {
					return err
				}

				switch keyd108ca1e39974529 {
				case "person":
					if dec.IsNull() {
						m.Person = nil
					} else if !dec.IsObjectOpen() {
						return errors.NewParseError(dec.Char(), dec.Cursor())
					} else {
						if dec.IsObjectClose() {
							m.Person = nil
						} else {
							if m.Person == nil {
								m.Person = new(CBPerson)
							}

							for objdb3f9883441f9783 := 1; objdb3f9883441f9783 > 0; {
								keya41ad111d1d89162, err := dec.NextKey()
								if err != nil {
									return err
								}

								switch keya41ad111d1d89162 {
								case "name":
									if dec.IsNull() {
										m.Person.Name = nil
									} else if !dec.IsObjectOpen() {
										return errors.NewParseError(dec.Char(), dec.Cursor())
									} else {
										if dec.IsObjectClose() {
											m.Person.Name = nil
										} else {
											if m.Person.Name == nil {
												m.Person.Name = new(CBName)
											}

											for obj18daf4578b33f100 := 1; obj18daf4578b33f100 > 0; {
												keyb2479e0673363dbb, err := dec.NextKey()
												if err != nil {
													return err
												}

												switch keyb2479e0673363dbb {
												case "fullName":
													valuec2d380bbadd0b8ef, err := dec.DecodeString()
													if err != nil {
														return err
													}

													m.Person.Name.FullName = valuec2d380bbadd0b8ef

												default:
													if err := dec.SkipValue(); err != nil {
														return err
													}
												}
												if dec.IsObjectClose() {
													obj18daf4578b33f100--
												}
											}
										}
									}

								case "github":
									if dec.IsNull() {
										m.Person.Github = nil
									} else if !dec.IsObjectOpen() {
										return errors.NewParseError(dec.Char(), dec.Cursor())
									} else {
										if dec.IsObjectClose() {
											m.Person.Github = nil
										} else {
											if m.Person.Github == nil {
												m.Person.Github = new(CBGithub)
											}

											for obj7dcb74d7e0b9c916 := 1; obj7dcb74d7e0b9c916 > 0; {
												key12fad66b0fa74d87, err := dec.NextKey()
												if err != nil {
													return err
												}

												switch key12fad66b0fa74d87 {
												case "followers":
													value97c7dcec08e3f545, err := dec.DecodeInt()
													if err != nil {
														return err
													}

													m.Person.Github.Followers = value97c7dcec08e3f545

												default:
													if err := dec.SkipValue(); err != nil {
														return err
													}
												}
												if dec.IsObjectClose() {
													obj7dcb74d7e0b9c916--
												}
											}
										}
									}

								case "gravatar":
									if dec.IsNull() {
										m.Person.Gravatar = nil
									} else if !dec.IsObjectOpen() {
										return errors.NewParseError(dec.Char(), dec.Cursor())
									} else {
										if dec.IsObjectClose() {
											m.Person.Gravatar = nil
										} else {
											if m.Person.Gravatar == nil {
												m.Person.Gravatar = new(CBGravatar)
											}

											for obj42edb89ddf1eb925 := 1; obj42edb89ddf1eb925 > 0; {
												keyd832c9079b14b3d6, err := dec.NextKey()
												if err != nil {
													return err
												}

												switch keyd832c9079b14b3d6 {
												case "avatars":
													if char := dec.NextChar(); char == 'n' {
														if err := dec.AssetNull(); err != nil {
															return err
														}

														m.Person.Gravatar.Avatars = nil
													} else if char != '[' {
														return errors.NewParseError(dec.Char(), dec.Cursor())
													} else {
														dec.Next()
														if dec.IsArrayClose() {
															m.Person.Gravatar.Avatars = nil
														} else {
															if m.Person.Gravatar.Avatars == nil {
																m.Person.Gravatar.Avatars = make(Avatars, 0, 8)
															}

															for arrayf6e34ae3e13fd8aa := 1; arrayf6e34ae3e13fd8aa > 0; {
																var value1de159b2673efb39 *CBAvatar
																if dec.IsNull() {
																	value1de159b2673efb39 = nil
																} else {
																	value1de159b2673efb39 = new(CBAvatar)

																	if char := dec.NextChar(); char == 'n' {
																		if err := dec.AssetNull(); err != nil {
																			return err
																		}

																	} else if char != '{' {
																		return errors.NewParseError(dec.Char(), dec.Cursor())
																	} else {
																		dec.Next()
																		if dec.IsObjectClose() {
																			return nil
																		} else {
																			for obj3cb97675310b4fb4 := 1; obj3cb97675310b4fb4 > 0; {
																				keyad9ed6979c62459b, err := dec.NextKey()
																				if err != nil {
																					return err
																				}

																				switch keyad9ed6979c62459b {
																				case "url":
																					valueeb83904df155de24, err := dec.DecodeString()
																					if err != nil {
																						return err
																					}

																					value1de159b2673efb39.Url = valueeb83904df155de24

																				default:
																					if err := dec.SkipValue(); err != nil {
																						return err
																					}
																				}
																				if dec.IsObjectClose() {
																					obj3cb97675310b4fb4--
																				}
																			}
																		}
																	}
																}
																if value1de159b2673efb39 != nil {
																	m.Person.Gravatar.Avatars = append(m.Person.Gravatar.Avatars, value1de159b2673efb39)
																}
																if dec.IsArrayClose() {
																	arrayf6e34ae3e13fd8aa--
																}
															}
														}
													}

												default:
													if err := dec.SkipValue(); err != nil {
														return err
													}
												}
												if dec.IsObjectClose() {
													obj42edb89ddf1eb925--
												}
											}
										}
									}

								default:
									if err := dec.SkipValue(); err != nil {
										return err
									}
								}
								if dec.IsObjectClose() {
									objdb3f9883441f9783--
								}
							}
						}
					}

				case "company":
					valuecb8ee83fd7af52fa, err := dec.DecodeString()
					if err != nil {
						return err
					}

					m.Company = valuecb8ee83fd7af52fa

				default:
					if err := dec.SkipValue(); err != nil {
						return err
					}
				}
				if dec.IsObjectClose() {
					obja785f5615ce785ac--
				}
			}
		}
	}
	dec.Release()

	return nil
}

func (s *SmallPayload) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()

	enc.WriteByte('{')
	enc.EncodeKeyInt("st", s.St)

	enc.EncodeKeyInt("sid", s.Sid)

	enc.EncodeKeyString("tt", s.Tt)

	enc.EncodeKeyInt("gr", s.Gr)

	enc.EncodeKeyString("uuid", s.Uuid)

	enc.EncodeKeyString("ip", s.Ip)

	enc.EncodeKeyString("ua", s.Ua)

	enc.EncodeKeyInt("tz", s.Tz)

	enc.EncodeKeyInt("v", s.V)

	enc.WriteByte('}')
	data := enc.Bytes()
	enc.Release()

	return data, nil
}

func (s *SmallPayload) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder()
	dec.SetUnsafeData(data)

	if char := dec.NextChar(); char == 'n' {
		if err := dec.AssetNull(); err != nil {
			return err
		}

	} else if char != '{' {
		return errors.NewParseError(dec.Char(), dec.Cursor())
	} else {
		dec.Next()
		if dec.IsObjectClose() {
			return nil
		} else {
			for objdd98bc4457ab9f11 := 1; objdd98bc4457ab9f11 > 0; {
				keyab80752ada368a12, err := dec.NextKey()
				if err != nil {
					return err
				}

				switch keyab80752ada368a12 {
				case "st":
					value5e7e006c87b5a679, err := dec.DecodeInt()
					if err != nil {
						return err
					}

					s.St = value5e7e006c87b5a679

				case "sid":
					value59e29df9e1e2d024, err := dec.DecodeInt()
					if err != nil {
						return err
					}

					s.Sid = value59e29df9e1e2d024

				case "tt":
					valuefdc571957bf549ae, err := dec.DecodeString()
					if err != nil {
						return err
					}

					s.Tt = valuefdc571957bf549ae

				case "gr":
					value7ff9a1f646942dae, err := dec.DecodeInt()
					if err != nil {
						return err
					}

					s.Gr = value7ff9a1f646942dae

				case "uuid":
					value4e42da0455eefb46, err := dec.DecodeString()
					if err != nil {
						return err
					}

					s.Uuid = value4e42da0455eefb46

				case "ip":
					value0f865cd07f1e2f7b, err := dec.DecodeString()
					if err != nil {
						return err
					}

					s.Ip = value0f865cd07f1e2f7b

				case "ua":
					valued8d285178a783aa8, err := dec.DecodeString()
					if err != nil {
						return err
					}

					s.Ua = valued8d285178a783aa8

				case "tz":
					valuefc23ae9ae32134ff, err := dec.DecodeInt()
					if err != nil {
						return err
					}

					s.Tz = valuefc23ae9ae32134ff

				case "v":
					value4bab422f130861dd, err := dec.DecodeInt()
					if err != nil {
						return err
					}

					s.V = value4bab422f130861dd

				default:
					if err := dec.SkipValue(); err != nil {
						return err
					}
				}
				if dec.IsObjectClose() {
					objdd98bc4457ab9f11--
				}
			}
		}
	}
	dec.Release()

	return nil
}
