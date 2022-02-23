/*
 * Copyright (C) 2019 Yunify, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this work except in compliance with the License.
 * You may obtain a copy of the License in the LICENSE file, or at:
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * 1. 支持select
 * 2. 支持where
 * 3. 支持简单表达式
 * 4. 支持action模板
 */

grammar TDTL;

// 1. Tokens & KeyWord
// 1.1 KeyWord
INSERT:                 I N S E R T;
INTO:                   I N T O;
AS:                     STUFF A S STUFF;
AND:                    STUFF A N D STUFF;
CASE:                   STUFF C A S E STUFF;
ELSE:                   STUFF E L S E STUFF;
END:                    STUFF E N D STUFF;
EQ:                     E Q     | '=';
FROM:                   STUFF F R O M STUFF;
GT:                     G T     | '>';
GTE:                    G T E   | '>' '=';
LT:                     L T     | '<';
LTE:                    L T E   | '<' '=';
NE:                     N E     | '!' '=' | '<' '>';
NOT:                    N O T   | '!';
NULL:                   N U L L;
OR:                     STUFF O R STUFF;
SELECT:                 S E L E C T STUFF;
THEN:                   STUFF T H E N STUFF;
WHERE:                  STUFF W H E R E STUFF;
WHEN:                   STUFF W H E N STUFF;


// 1.2 Token
MUL:                '*';
DIV:                '/';
MOD:                '%';
ADD:                '+';
SUB:                '-';
DOT:                '.';
TRUE:               T R U E;
FALSE:              F A L S E;
INDENTIFIER:        [a-zA-Z0-9_#][a-zA-Z_\-#$@0-9]*;
NUMBER:             '0' | [1-9][0-9]* ;
INTEGER:            ('+' | '-')? NUMBER;
FLOAT:              ('+' | '-')? (NUMBER+ DOT NUMBER+ |  NUMBER+ DOT | DOT NUMBER+);
TOPICITEM:          [a-zA-Z_/\-#$@0-9]+;
PATHITEM:           TOPICITEM (ARRAYITEM)? (DOT TOPICITEM (ARRAYITEM)?)*;
ARRAYITEM:          '[' NUMBER ']' | '[' '#' ']';
STRING:             '\'' (~'\'' | '\'\'')* '\'';
WHITESPACE:         [ \r\n\t]+ -> skip;




// 2. Rules
root
    : INSERT INTO target SELECT fields EOF;

target
    : INDENTIFIER
    ;

// 2.1 Select
fields
    : field_elem (',' field_elem)*
    ;


field_elem
    : field_elem_with_as                            # FieldElemAs
    | sourceEntity '.' asterisk                     # FieldElemSource
    | expr                                          # FieldElemExpr
    ;

field_elem_with_as
    : expr AS target_name                            # TargetAsElem
    ;



filter
    : filter_condition
    ;

filter_condition
    : filter_condition_or (AND filter_condition_or)*
    ;

filter_condition_or
    : filter_condition_not (OR filter_condition_not)*
    ;

filter_condition_not
    : NOT? expr
    ;

// 2.4 expr
expr
   : constant                                       # Braces
   | '(' expr ')'                                   # Braces
   | expr op=('*'|'/'|'%') expr                     # Binary
   | expr op=('+'|'-') expr                         # Binary
   | expr op=(EQ | GT | LT | GTE | LTE | NE) expr   # Binary
   | call_expr                                      # Function
   | switch_stmt                                    # Switch
   ;

// 2.1 entity
/*
source_stmt
    : sourceEntity propertyEntity                   # ExprElemSource
    ;
*/

sourceEntity
    :INDENTIFIER
    ;

propertyEntity
    :('.' INDENTIFIER)+
    ;

constant
    : TRUE                                           # Boolean
    | FALSE                                          # Boolean
    | NUMBER                                         # Integer
    | INTEGER                                        # Integer
    | FLOAT                                          # Float
    | STRING                                         # String
    | xpath_name                                     # Source
    ;


switch_stmt
    : CASE expr
    WHEN expr THEN expr (WHEN expr THEN expr)*
    (ELSE expr)?
    ;

call_expr
    : key=INDENTIFIER '(' (expr (',' expr)*)? ')';

/*
CASE v WHEN t[1] THEN r[1]
  WHEN t[2] THEN r[2] ...
  WHEN t[n] THEN r[n]
  ELSE r[e] END
*/

// 2.5 json
asterisk: '*';

xpath_name
        : dotnotation
        | '"' + (dotnotation) + '"'
        ;

target_name
        : dotnotation
        | '"' + (dotnotation) + '"'
        ;


dotnotation
    : INDENTIFIER
    | PATHITEM
    ;

identifierWithTOPICITEM
    :  PATHITEM '[' ']'
    | PATHITEM '[' NUMBER ']'
    | PATHITEM '[' '#'  ']'
    | PATHITEM
    |  FLOAT
    ;

identifierWithQualifier : INDENTIFIER '[]'
                        | INDENTIFIER '[' NUMBER ']'
                        | INDENTIFIER '[#]'
                        | INDENTIFIER
                        ;



fragment A: [aA];
fragment B: [bB];
fragment C: [cC];
fragment D: [dD];
fragment E: [eE];
fragment F: [fF];
fragment G: [gG];
fragment H: [hH];
fragment I: [iI];
fragment J: [jJ];
fragment K: [kK];
fragment L: [lL];
fragment M: [mM];
fragment N: [nN];
fragment O: [oO];
fragment P: [pP];
fragment Q: [qQ];
fragment R: [rR];
fragment S: [sS];
fragment T: [tT];
fragment U: [uU];
fragment V: [vV];
fragment W: [wW];
fragment X: [xX];
fragment Y: [yY];
fragment Z: [zZ];
fragment STUFF: [ \r\n\t]+;
