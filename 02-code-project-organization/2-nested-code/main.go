package main

import "errors"

// ネストが深い
func join1(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

// 左がハッピーパス　右がエラーケース（バッドパス）
// 逆にならないうようにした方が可読性が上がる 例: if s1 != "" { 通常の処理 } は書き換えた方が良い
func join2(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil
}

func concatenate(s1, s2 string) (string, error) {
	return "", nil
}
